package connector

import (
	connectorpb "github.com/nutanix/kps-connector-go-sdk/connector/v1"
	"github.com/nutanix/kps-connector-go-sdk/events"
)

var (
	transportPublishFailedAlert     = events.NewAlert("transportPublishFailed", "failed to publish message on transport", connectorpb.Severity_SEVERITY_CRITICAL, connectorpb.State_STATE_FAILED)
	transportSubscribeFailedAlert   = events.NewAlert("transportSubscribeFailed", "failed to subscribe to transport", connectorpb.Severity_SEVERITY_CRITICAL, connectorpb.State_STATE_FAILED)
	transportUnsubscribeFailedAlert = events.NewAlert("transportUnsubscribeFailed", "failed to unsubscribe from transport", connectorpb.Severity_SEVERITY_CRITICAL, connectorpb.State_STATE_FAILED)

	streamStartedStatus   = events.NewStatus("streamStarted", "stream has successfully started", connectorpb.State_STATE_PROVISIONED)
	streamHealthyStatus   = events.NewStatus("streamHealthy", "stream is healthy", connectorpb.State_STATE_HEALTHY)
	streamUnhealthyStatus = events.NewStatus("streamUnhealthy", "stream is unhealthy", connectorpb.State_STATE_UNHEALTHY)
)

func (d *Connector) initEventRegistry() {
	d.RegisterAlert(transportPublishFailedAlert)
	d.RegisterAlert(transportSubscribeFailedAlert)
	d.RegisterAlert(transportUnsubscribeFailedAlert)
	d.RegisterStatus(streamStartedStatus)
	d.RegisterStatus(streamHealthyStatus)
	d.RegisterStatus(streamUnhealthyStatus)
}
