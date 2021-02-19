package transport

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
	connectorpb "github.com/nutanix/kps-connector-go-sdk/connector/v1"
	"github.com/nutanix/kps-connector-go-sdk/internal"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

type cfg struct {
	NatsBroker          string
	Name                string
	pushgatewayEndpoint string
}

var (
	transportCfg = &cfg{
		NatsBroker:          os.Getenv("NATS_BROKER"),
		Name:                os.Getenv("NATS_NAME"),
		pushgatewayEndpoint: os.Getenv("PUSH_GW"),
	}
	transportConnectErrorCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: fmt.Sprintf("transport_connect_errors"),
		Help: "Number of NATS connect errors encountered",
	})
	transportPublishErrorCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: fmt.Sprintf("transport_publish_errors"),
		Help: "Number of NATS messages that encountered errors",
	})
	statsRegistry = prometheus.NewRegistry()

	singleton Client
	once      internal.Once
)

func init() {
	statsRegistry.MustRegister(transportConnectErrorCounter, transportPublishErrorCounter)

	// Start the pushgateway pusher go routine which periodically pushes
	// prometheus metrics of transport to pushgateway
	go func() {
		stopMetricsPusher := make(chan bool)
		defer close(stopMetricsPusher)

		metricsPushTicker := time.NewTicker(1 * time.Minute)
		defer metricsPushTicker.Stop()

		for {
			select {
			case <-stopMetricsPusher:
				return
			case <-metricsPushTicker.C:
				_ = push.New(transportCfg.pushgatewayEndpoint, "connector_transport_metrics_job").Gatherer(statsRegistry).Push()
			}
		}
	}()
}

// Message defines the data structure of the messages conveyed by the transport
type Message struct {
	Payload []byte `json:"payload"`
}

// MessageHandler defines the function signature for the callback function in a Subscribe call
type MessageHandler func(msg *Message)

// Client describes the publish / subscribe interface of the transport client
type Client interface {
	// Publish publishes the message onto the provided channel
	Publish(channel string, msg Message) error
	// Subscribe subscribes all future messages on the channel and registers a callback
	Subscribe(channel string, callback MessageHandler) (Subscription, error)
}

// Subscription describes the interface of the subscription object
type Subscription interface {
	// Unsubscribe unsubscribes the connection
	Unsubscribe() error
	// Channel returns the channel the subscription belongs to
	Channel() string
}

type natsSubscription struct {
	*nats.Subscription
}

// Unsubscribe unsubscribes the connection
func (sub *natsSubscription) Unsubscribe() error {
	return sub.Subscription.Unsubscribe()
}

// Channel returns the channel the subscription belongs to
func (sub *natsSubscription) Channel() string {
	return sub.Subject
}

// NewTransportClient returns a client for publishing and subscribing to datastreams from data pipelines
func NewTransportClient() (Client, error) {
	err := once.TryDo(func() error {
		natsClient, err := newNatsClient(transportCfg.NatsBroker, transportCfg.Name)
		if err != nil {
			glog.Errorf("Failed to connect to Transport Broker: %s", err.Error())
			return err
		}
		singleton = natsTransportClient(natsClient)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return singleton, nil
}

type natsClient struct {
	conn *nats.Conn
	url  string
}

var _ Client = (*natsClient)(nil)

// create the underlying nats.Conn object
func newNatsClient(url string, id string) (*nats.Conn, error) {
	return nats.Connect(url, nats.Name(id),
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			fmt.Printf("Got disconnected! Reason: %q\n", err)
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			fmt.Printf("Got reconnected to %v!\n", nc.ConnectedUrl())
		}),
		nats.ClosedHandler(func(nc *nats.Conn) {
			fmt.Printf("Connection closed. Reason: %q\n", nc.LastError())
		}))
}

// natsTransportClient wraps a nats.Conn into a transport.Client object
func natsTransportClient(client *nats.Conn) Client {
	natsClientInst := &natsClient{
		conn: client,
		url:  client.ConnectedAddr(),
	}

	return natsClientInst
}

// Publish publishes the message onto the provided channel
func (client *natsClient) Publish(subject string, msg Message) error {
	tMsg := connectorpb.TransportMessage{
		Timestamp: time.Now().UnixNano(),
		Payload:   [][]byte{msg.Payload},
	}

	data, err := proto.Marshal(&tMsg)
	if err != nil {
		transportPublishErrorCounter.Inc()
		return err
	}

	err = client.conn.Publish(subject, data)
	if err != nil {
		transportPublishErrorCounter.Inc()
		return err
	}

	return nil
}

// Subscribe subscribes all future messages on the channel and registers a callback
func (client *natsClient) Subscribe(subject string, cb MessageHandler) (Subscription, error) {
	natsSub, err := client.conn.Subscribe(subject, client.natsMsgHandler(cb))
	if err != nil {
		return nil, err
	}
	return &natsSubscription{Subscription: natsSub}, nil
}

func (client *natsClient) natsMsgHandler(handler MessageHandler) nats.MsgHandler {
	return func(msg *nats.Msg) {
		var tMsg connectorpb.TransportMessage
		err := proto.Unmarshal(msg.Data, &tMsg)
		if err != nil {
			log.Printf("unable to unmarshal data from %s", msg.Subject)
		}
		hMsg := &Message{
			Payload: tMsg.Payload[0],
		}
		handler(hMsg)
	}
}
