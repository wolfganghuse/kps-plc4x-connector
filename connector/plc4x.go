package connector

import (
	"context"
	"log"
	"time"

	"github.com/nutanix/kps-connector-go-sdk/transport"

	"github.com/apache/plc4x/plc4go/pkg/plc4go"
	"github.com/apache/plc4x/plc4go/pkg/plc4go/drivers"
	"github.com/apache/plc4x/plc4go/pkg/plc4go/model"
	"github.com/apache/plc4x/plc4go/pkg/plc4go/transports"
)

type Address struct {
	Name    string
	Address string
}

type streamMetadata struct {
	Plc       string
	Addresses []Address

	
	//PollingIntervall int  // removed for testing
}

// mapToStreamMetadata translates the stream metadata into the corresponding streamMetadata struct
func mapToStreamMetadata(metadata map[string]interface{}) *streamMetadata {
	plc := metadata["plc"].(string)
	addressObjs, ok := metadata["addresses"].([]interface{})
	if !ok {
		// bail out
	}
	addresses := make([]Address, 0, len(addressObjs))
	for _, obj := range addressObjs {
		addressMap, ok := obj.(map[string]string)
		if !ok {
			// bail out
		}
		addressName := addressMap["name"]
		log.Printf("add Name: %s", addressName)
		addressAddress := addressMap["address"]
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


	if rrr.Response.GetResponseCode("field1") != model.PlcResponseCode_OK {
		log.Printf("error an non-ok return code: %s", rrr.Response.GetResponseCode("field1").GetName())
		return nil, nil
	}

	value := rrr.Response.GetValue("field1")
	log.Printf("Returned Value: %s", value.GetString())
	
	return []byte(value.GetString()), nil // ???? TODO Typ-Conversion to Bytes missing
}

// subscribe wraps the logic to connect or subscribe to the corresponding stream
// from the relevant client or service
func (c *consumer) subscribe(ctx context.Context, metadata *streamMetadata) error {

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
		//rrb.AddItem("field1","holding-register:4:INT")
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
