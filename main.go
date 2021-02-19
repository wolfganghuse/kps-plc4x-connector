package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	connectorpb "github.com/nutanix/kps-connector-go-sdk/connector/v1"
	"github.com/nutanix/kps-connector-go-template/connector"
	"google.golang.org/grpc"
)

const (
	// this is the port on which the connector instance will be listening to for the GRPC calls
	connectorInstanceServerPort = 8000
)

func main() {
	log.Printf("%s - Start", connector.ConnectorCfg.Name)

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", connectorInstanceServerPort))
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	log.Printf("successfully set up listener on %d\n", connectorInstanceServerPort)

	var opts []grpc.ServerOption

	srv := connector.NewConnector()

	grpcServer := grpc.NewServer(opts...)
	connectorpb.RegisterConnectorServiceServer(grpcServer, srv)

	log.Printf("starting to serve grpc server on %s\n", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to run the grpc server: %s", err)
	}
}

func init() {
	// initialize with default values as static config
	flag.StringVar(&connector.ConnectorCfg.ID,
		"connectorInstanceID",
		os.Getenv("CONNECTOR_INSTANCE_ID"),
		"UUID of the connector instance")

	flag.Parse()
}
