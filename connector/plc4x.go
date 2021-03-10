package connector

import (
	"context"
	"encoding/json"
	"log"
	"time"
	"flag"

	"github.com/nutanix/kps-connector-go-sdk/transport"

	"github.com/apache/plc4x/plc4go/pkg/plc4go"
	"github.com/apache/plc4x/plc4go/pkg/plc4go/drivers"
	"github.com/apache/plc4x/plc4go/pkg/plc4go/model"
	"github.com/apache/plc4x/plc4go/pkg/plc4go/transports"
	
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)


type Address struct {
	Name    string
	Address string
}

type rvalue struct{
	Field string `json:"field"`
	Value string `json:"value"`
}


type streamMetadata struct {
	Plc       string
	Addresses []Address

	
	//PollingIntervall int  // removed for testing
}

var (
	netBytes         = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "network_bytes_total",
		Help: "Number of bytes seen on the network.",
	})

	addr              = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")
)

// mapToStreamMetadata translates the stream metadata into the corresponding streamMetadata struct
func mapToStreamMetadata(metadata map[string]interface{}) *streamMetadata {
	plc := metadata["plc"].(string)
	addressObjs, ok := metadata["addresses"].([]interface{})
	if !ok {
		log.Printf("unable to type infer Addresses: expected []interface found %T", metadata["addresses"])
	}
	addresses := make([]Address, 0, len(addressObjs))
	for _, obj := range addressObjs {
		addressMap, ok := obj.(map[string]interface{})
		if !ok {
			// bail out
		}
		addressName, ok := addressMap["name"].(string)
		if !ok {
			// bail out
		}
		log.Printf("add Name: %s", addressName)
		addressAddress, ok := addressMap["address"].(string)
		if !ok {
			// bail out
		}
		log.Printf("add Address: %s", addressAddress)
		
		addr := Address{
			Name: addressName,
			Address: addressAddress,
		}
		addresses = append(addresses, addr)
	}


	return &streamMetadata{
		Plc: plc,
		Addresses: addresses,		
	}
}

type consumer struct {
	transport  transport.Client
	metadata   *streamMetadata
	connection *plc4go.PlcConnection    //?????
	rr         model.PlcReadRequest // ???? ReadRequest}
}

// producer consumes the data from the relevant client or service and publishes them to KPS data pipelines
func newConsumer() *consumer {
	// TODO: Add the relevant clients and fields
	return &consumer{}
}

// nextMsg wraps the logic for consuming iteratively a transport.Message
// from the relevant client or service
func (c *consumer) nextMsg() ([]byte, error) {


	netBytes.Add(float64(1))

	// manual slow down Polling
	time.Sleep(5 * time.Second )
	rrc := c.rr.Execute()
	log.Printf("nexmsg")
	// Wait for the response to finish
	rrr := <-rrc
	if rrr.Err != nil {
		log.Printf("error executing read-request: %s", rrr.Err.Error())
		return nil, nil
	}

	// Do something with the response
	
	rvalues := make([]rvalue,0)
	
	for _, fieldname := range rrr.Response.GetFieldNames() {
		log.Printf("processing field: %s", fieldname)
		if rrr.Response.GetResponseCode(fieldname) != model.PlcResponseCode_OK {
			log.Printf("error an non-ok return code: %s", rrr.Response.GetResponseCode(fieldname).GetName())
			return nil, nil
		}
	
		value := rrr.Response.GetValue(fieldname)
		log.Printf("Returned Value: %s", value.GetString())
//
		
		toMarshal := rvalue{
			Field: fieldname,
			Value: value.GetString(),
		}
		
//
		rvalues = append(rvalues,toMarshal)

		
	}
	
	
	return json.Marshal(rvalues)

}

// subscribe wraps the logic to connect or subscribe to the corresponding stream
// from the relevant client or service
func (c *consumer) subscribe(ctx context.Context, metadata *streamMetadata) error {

	prometheus.MustRegister(netBytes)

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe(*addr, nil))
	}()

	driverManager := plc4go.NewPlcDriverManager()
	transports.RegisterTcpTransport(driverManager)
	drivers.RegisterModbusDriver(driverManager)

	// Get a connection to a remote PLC
	crc := driverManager.GetConnection(metadata.Plc)
	 // Wait for the driver to connect (or not)
	connectionResult := <-crc
	if connectionResult.Err != nil {
		log.Printf("error connecting to PLC: %s", connectionResult.Err.Error())
		return nil
	}
	connection := connectionResult.Connection

	// Make sure the connection is closed at the end
	defer connection.Close()

	// Prepare a read-request
	rrb := connection.ReadRequestBuilder()
	//rrb.AddItem("field1","holding-register:4:INT")
	for _, address := range metadata.Addresses {
		log.Printf("added name: %s", address.Name)
		log.Printf("added address: %s", address.Address)
		rrb.AddItem(address.Name, address.Address)
	}
	
	readRequest, err := rrb.Build()
	if err != nil {
		log.Printf("error preparing read-request: %s", connectionResult.Err.Error())
		return nil
	}
	c.rr = readRequest
	return nil
}

// producer produces data received from KPS data pipelines to the relevant client
type producer struct {
	// TODO: Add the relevant client and fields
}

func newProducer() *producer {
	// TODO: Add the relevant client and fields
	return &producer{}
}

func (p *producer) connect(ctx context.Context, metadata *streamMetadata) error {
	// TODO: Add the code for connecting to the relevant client or service for publishing to it
	return nil
}

// subscribeMsgHandler is a callback function that wraps the logic for producing a transport.Message
// from the data pipelines into the relevant client or service
func (*producer) subscribeMsgHandler(message *transport.Message) {
	// TODO: Add code for producing data into the relevant client
	log.Printf("msg received: %s", string(message.Payload))
}
