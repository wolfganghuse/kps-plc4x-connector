package connector

import (
	"context"
	"log"
	"reflect"

	"github.com/nutanix/kps-connector-go-sdk/transport"

	"github.com/wolfganghuse/kps-plc4x-connector/plc4x/plc4go/internal/plc4go/modbus"
	"github.com/wolfganghuse/kps-plc4x-connector/plc4x/plc4go/internal/plc4go/spi/transports/tcp"
	"github.com/wolfganghuse/kps-plc4x-connector/plc4x/plc4go/pkg/plc4go"
	"github.com/wolfganghuse/kps-plc4x-connector/plc4x/plc4go/pkg/plc4go/model"
)

type streamMetadata struct {
	Plc string
	Addresses  struct {
		Name string
		Address string
	}
	PollingIntervall int
}

// mapToStreamMetadata translates the stream metadata into the corresponding streamMetadata struct
func mapToStreamMetadata(metadata map[string]interface{}) *streamMetadata {
	plc := metadata["plc"].(string)
	// ???? How to iterate through the unkown size of a struct passed into Metadata?
	//addresses := metadata[""].(struct)
	pollingintervall := metadata["polling-intervall"].(string)
	return &streamMetadata{
		Plc:  plc,
		// Same here, need to assign the struct here
		// Addresses: {Name, Address}
		PollingIntervall: pollingintervall,
	}
}

type consumer struct {
	transport transport.Client
	metadata  *streamMetadata
	connection   * //?????
	rr     chan * // ???? ReadRequest}

// producer consumes the data from the relevant client or service and publishes them to KPS data pipelines
func newConsumer() *consumer {
	// TODO: Add the relevant clients and fields
	return &consumer{}
}

// nextMsg wraps the logic for consuming iteratively a transport.Message
// from the relevant client or service
func (c *consumer) nextMsg() ([]byte, error) {
	
	rrc := c.rr.Execute()

	// Wait for the response to finish
	rrr := <-rrc
	if rrr.Err != nil {
		log.Printf("error executing read-request: %s", rrr.Err.Error())
		return
	}

	// Do something with the response
	if rrr.Response.GetResponseCode("field") != model.PlcResponseCode_OK {
		log.Printf("error an non-ok return code: %s", rrr.Response.GetResponseCode("field").GetName())
		return
	}

	value := rrr.Response.GetValue("field")
	return value.Data, nil
}

// subscribe wraps the logic to connect or subscribe to the corresponding stream
// from the relevant client or service
func (c *consumer) subscribe(ctx context.Context, metadata *streamMetadata) error {

	driverManager := plc4go.NewPlcDriverManager()
	driverManager.RegisterDriver(modbus.NewModbusDriver())
	driverManager.RegisterTransport(tcp.NewTcpTransport())

	// Get a connection to a remote PLC
	crc := driverManager.GetConnection("modbus:tcp://192.168.178.183")

	// Wait for the driver to connect (or not)
	connectionResult := <-crc
	if connectionResult.Err != nil {
		log.Printf("error connecting to PLC: %s", connectionResult.Err.Error())
		return
	}
	connection := connectionResult.Connection

	// Make sure the connection is closed at the end
	defer connection.Close()

	// Prepare a read-request
	rrb := connection.ReadRequestBuilder()
	rrb.AddItem("field", "holding-register:26:REAL")
	readRequest, err := rrb.Build()
	if err != nil {
		log.Printf("error preparing read-request: %s", connectionResult.Err.Error())
		return
	}

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
