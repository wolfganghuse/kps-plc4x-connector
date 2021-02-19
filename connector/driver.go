package connector

import (
	"context"
	"fmt"
	"log"
	"sync"

	connectorpb "github.com/nutanix/kps-connector-go-sdk/connector/v1"
	"github.com/nutanix/kps-connector-go-sdk/events"
	"github.com/nutanix/kps-connector-go-sdk/transport"
)

// Connector implements the ConnectorService gRPC service
type Connector struct {
	mtx              sync.RWMutex
	id               string
	streams          []*connectorpb.Stream
	activeInStreams  map[string]context.CancelFunc
	activeOutStreams map[string]transport.Subscription

	// Registry implements the `GetEvents` method
	*events.Registry
	connectorpb.UnsafeConnectorServiceServer
}

var _ connectorpb.ConnectorServiceServer = (*Connector)(nil)

// NewConnector is a constructor for the Connector object
func NewConnector() *Connector {
	registry := events.NewRegistry()
	d := &Connector{
		id:               ConnectorCfg.ID,
		streams:          make([]*connectorpb.Stream, 0),
		activeInStreams:  make(map[string]context.CancelFunc),
		activeOutStreams: make(map[string]transport.Subscription),
		Registry:         registry,
	}

	d.initEventRegistry()
	return d
}

// GetPayload returns all payloads given a payload kind:
//   - If payload kind is set to STREAM, it should return all available streams (both subscribed and not subscribed)
//   - If payload kind is set to CONFIG, it should return the current dynamic config in use.
func (d *Connector) GetPayload(ctx context.Context, req *connectorpb.GetPayloadRequest) (*connectorpb.GetPayloadResponse, error) {
	if req.GetConnectorId() != d.id {
		err := fmt.Errorf("wrong Connector id")
		resp := &connectorpb.GetPayloadResponse{Status: &connectorpb.ResponseStatus{Code: connectorpb.ResponseCode_RESPONSE_CODE_INVALID_ARGUMENT, Message: err.Error()}}
		return resp, err
	}

	switch req.GetKind() {
	case connectorpb.PayloadKind_PAYLOAD_KIND_STREAM:
		return d.getStreams(ctx)
	case connectorpb.PayloadKind_PAYLOAD_KIND_CONFIG:
		return d.getConfig(ctx)
	}

	err := fmt.Errorf("unknown payload kind")
	resp := &connectorpb.GetPayloadResponse{Status: &connectorpb.ResponseStatus{Code: connectorpb.ResponseCode_RESPONSE_CODE_INVALID_ARGUMENT, Message: err.Error()}}
	return resp, err
}

// SetPayload updates the payloads that connector needs to process:
//   - If payload kind is set to STREAM, it should reset the topics being used by subscribing to the streams from the request and unsubscribing the old ones
//   - If payload kind is set to CONFIG, it should update the current dynamic config
func (d *Connector) SetPayload(ctx context.Context, req *connectorpb.SetPayloadRequest) (*connectorpb.SetPayloadResponse, error) {
	if req.GetConnectorId() != d.id {
		err := fmt.Errorf("wrong Connector id")
		resp := &connectorpb.SetPayloadResponse{Status: &connectorpb.ResponseStatus{Code: connectorpb.ResponseCode_RESPONSE_CODE_INVALID_ARGUMENT, Message: err.Error()}}
		return resp, err
	}

	payloads := req.Payloads

	configs := make([]*connectorpb.Config, 0)
	streams := make([]*connectorpb.Stream, 0)
	for _, payload := range payloads {
		if topic := payload.GetStream(); topic != nil {
			streams = append(streams, topic)
		} else if config := payload.GetConfig(); config != nil {
			configs = append(configs, config)
		} else {
			log.Printf("payload is neither a config nor a stream: %+v", payload)
		}
	}

	d.updateConfig(ctx, configs)

	if err := d.setStreams(ctx, streams); err != nil {
		resp := &connectorpb.SetPayloadResponse{Status: &connectorpb.ResponseStatus{Code: connectorpb.ResponseCode_RESPONSE_CODE_INTERNAL, Message: err.Error()}}
		return resp, err
	}

	return &connectorpb.SetPayloadResponse{Status: &connectorpb.ResponseStatus{Code: connectorpb.ResponseCode_RESPONSE_CODE_OK}}, nil
}
