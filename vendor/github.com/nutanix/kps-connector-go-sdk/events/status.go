package events

import (
	"fmt"

	connectorpb "github.com/nutanix/kps-connector-go-sdk/connector/v1"
)

// StatusOpts defines the type for the functional options for publishing status
type StatusOpts func(*statusInst)

// Status is the interface for updating status
type Status interface {
	Publish(...StatusOpts) error
}

type statusImpl struct {
	name     string
	id       string
	message  string
	state    connectorpb.State
	registry *Registry
}

var _ Status = (*statusImpl)(nil)

// NewStatus creates a status object that can be registered with the registry
func NewStatus(name string, message string, state connectorpb.State) Status {
	return &statusImpl{
		name:    name,
		message: message,
		state:   state,
	}
}

type statusInst struct {
	status   *statusImpl
	streamID string
	metadata *EventMetadata
}

// StatusWithStreamID is for publishing a stream specific status
func StatusWithStreamID(streamID string) StatusOpts {
	return func(inst *statusInst) {
		inst.streamID = streamID
	}
}

// StatusWithEventMetadata is for publishing metadata along with Status
func StatusWithEventMetadata(metadata *EventMetadata) StatusOpts {
	return func(inst *statusInst) {
		inst.metadata = metadata
	}
}

// Publish publishes the status with the provided options
func (s *statusImpl) Publish(opts ...StatusOpts) error {
	if s.registry == nil {
		return fmt.Errorf("status not registered with the registry")
	}

	inst := &statusInst{
		status: s,
	}
	for _, opt := range opts {
		opt(inst)
	}
	statusEvent := &connectorpb.Status{
		Id:       s.name,
		StreamId: inst.streamID,
		Message:  s.message,
		State:    s.state,
	}

	if inst.metadata != nil {
		metadata, err := inst.metadata.toStruct()
		if err != nil {
			return err
		}
		statusEvent.Metadata = metadata
	}

	s.registry.addEvent(statusEvent)
	return nil
}

// String creates a stringified representation of the status
func (s *statusImpl) String() string {
	return fmt.Sprintf("[name: %s][message: %s][state: %s]", s.name, s.message, s.state)
}
