package events

import (
	"fmt"

	connectorpb "github.com/nutanix/kps-connector-go-sdk/connector/v1"
)

// AlertOpts defines the type for the functional options for publishing alerts
type AlertOpts func(*alertInst)

// Alert is the interface for raising alerts
type Alert interface {
	Publish(...AlertOpts) error
}

type alertImpl struct {
	name     string
	message  string
	severity connectorpb.Severity
	state    connectorpb.State
	registry *Registry
}

var _ Alert = (*alertImpl)(nil)

// NewAlert creates an alert object that can be registered with the registry
func NewAlert(name string, message string, severity connectorpb.Severity, state connectorpb.State) Alert {
	return &alertImpl{
		name:     name,
		message:  message,
		severity: severity,
		state:    state,
	}
}

type alertInst struct {
	alert    *alertImpl
	streamID string
	metadata *EventMetadata
}

// AlertWithStreamID is for publishing a stream specific alert
func AlertWithStreamID(streamID string) AlertOpts {
	return func(inst *alertInst) {
		inst.streamID = streamID
	}
}

// AlertWithEventMetadata is for publishing metadata along with Alert
func AlertWithEventMetadata(metadata *EventMetadata) AlertOpts {
	return func(inst *alertInst) {
		inst.metadata = metadata
	}
}

// Publish publishes the alert with the provided options
func (a *alertImpl) Publish(opts ...AlertOpts) error {
	if a.registry == nil {
		return fmt.Errorf("alert not registered with the registry")
	}

	inst := &alertInst{
		alert: a,
	}
	for _, opt := range opts {
		opt(inst)
	}

	alertEvent := &connectorpb.Alert{
		Id:       a.name,
		StreamId: inst.streamID,
		Message:  a.message,
		Severity: a.severity,
		State:    a.state,
	}

	if inst.metadata != nil {
		metadata, err := inst.metadata.toStruct()
		if err != nil {
			return err
		}
		alertEvent.Metadata = metadata
	}

	a.registry.addEvent(alertEvent)
	return nil
}

// String creates a stringified representation of the alert
func (a *alertImpl) String() string {
	return fmt.Sprintf("[name: %s][message: %s][severity: %s][state: %s]", a.name, a.message, a.severity, a.state)
}
