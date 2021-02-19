package connector

import (
	"context"
	"sync"

	connectorpb "github.com/nutanix/kps-connector-go-sdk/connector/v1"
	"google.golang.org/protobuf/types/known/structpb"
)

//Config type for all configuration data
type Config struct {
	Name string
	ID   string

	sync.RWMutex
	dynamicConfig map[string]interface{}
}

var (
	// ConnectorCfg stores the dynamic configuration of the Connector runtime
	ConnectorCfg = &Config{
		// TODO: Set the name for the Connector
		Name:          "TemplateConnector",
		dynamicConfig: make(map[string]interface{}),
	}
)

func (d *Connector) getConfig(context.Context) (*connectorpb.GetPayloadResponse, error) {
	ConnectorCfg.RLock()
	defer ConnectorCfg.RUnlock()
	payloads := make([]*connectorpb.Payload, 0)
	cfg, err := structpb.NewStruct(ConnectorCfg.dynamicConfig)
	if err != nil {
		status := &connectorpb.ResponseStatus{Code: connectorpb.ResponseCode_RESPONSE_CODE_INTERNAL, Message: err.Error()}
		return &connectorpb.GetPayloadResponse{Status: status}, err
	}

	payloadConfig := &connectorpb.Payload_Config{Config: &connectorpb.Config{Metadata: cfg}}
	payload := &connectorpb.Payload{Object: payloadConfig}
	payloads = append(payloads, payload)

	resp := &connectorpb.GetPayloadResponse{
		Status:   &connectorpb.ResponseStatus{Code: connectorpb.ResponseCode_RESPONSE_CODE_OK},
		Payloads: payloads,
	}

	return resp, nil
}

func (d *Connector) updateConfig(ctx context.Context, configs []*connectorpb.Config) {
	ConnectorCfg.Lock()
	defer ConnectorCfg.Unlock()
	for _, config := range configs {
		for k, v := range config.GetMetadata().AsMap() {
			ConnectorCfg.dynamicConfig[k] = v
		}
	}
}
