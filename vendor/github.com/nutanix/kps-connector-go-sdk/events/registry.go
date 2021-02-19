package events

import (
	"context"
	"sync"

	"github.com/golang/glog"
	connectorpb "github.com/nutanix/kps-connector-go-sdk/connector/v1"
)

// Registry enables registering events that need to be published. It also implements the GetEvents
// method required for fulfilling the data connector contract. This ensures that an embedded registry object
// provides everything a connector needs when it comes to handling events
type Registry struct {
	alerts   map[string]*alertImpl
	statuses map[string]*statusImpl
	events   map[string]interface{}
	rwLock   sync.RWMutex
}

// NewRegistry creates a new events registry for use in a data connector
func NewRegistry() *Registry {
	return &Registry{
		rwLock:   sync.RWMutex{},
		events:   make(map[string]interface{}),
		alerts:   make(map[string]*alertImpl),
		statuses: make(map[string]*statusImpl),
	}
}

// RegisterAlert registers an alert event in the registry
func (reg *Registry) RegisterAlert(alert Alert) {
	a := alert.(*alertImpl)
	reg.alerts[a.name] = a
	a.registry = reg
}

// RegisterStatus registers a status event in the registry
func (reg *Registry) RegisterStatus(status Status) {
	s := status.(*statusImpl)
	reg.statuses[s.name] = s
	s.registry = reg
}

// GetEvents implements the GetEvents method required for fulfilling data connector contract
func (reg *Registry) GetEvents(context.Context, *connectorpb.GetEventsRequest) (*connectorpb.GetEventsResponse, error) {
	events := reg.getAllConnectorEvents()
	resp := &connectorpb.GetEventsResponse{
		Status: &connectorpb.ResponseStatus{
			Code: connectorpb.ResponseCode_RESPONSE_CODE_OK,
		},
		EventPayloads: events,
	}
	return resp, nil
}

func (reg *Registry) addEvent(event interface{}) {
	reg.rwLock.Lock()
	defer reg.rwLock.Unlock()
	var key string
	switch event := event.(type) {
	case *connectorpb.Alert:
		key = event.Id + event.StreamId
	case *connectorpb.Status:
		key = event.Id + event.StreamId
	}

	reg.events[key] = event
	return
}

func (reg *Registry) getAllEvents() (events []interface{}) {
	reg.rwLock.RLock()
	defer reg.rwLock.RUnlock()

	for _, event := range reg.events {
		events = append(events, event)
	}
	return events
}

func (reg *Registry) logAllEvents() {
	reg.rwLock.RLock()
	defer reg.rwLock.RUnlock()

	for key, event := range reg.events {
		glog.Infof("[%s] event: => %#v", key, event)
	}
}

func (reg *Registry) clearCache() {
	reg.rwLock.Lock()
	defer reg.rwLock.Unlock()

	reg.events = make(map[string]interface{})

	return
}

func (reg *Registry) getAllConnectorEvents() []*connectorpb.EventPayload {
	events := reg.getAllEvents()

	eventPayloads := make([]*connectorpb.EventPayload, 0)
	for _, v := range events {
		switch event := v.(type) {
		case *connectorpb.Alert:
			alertEventPayload := &connectorpb.EventPayload_Alert{Alert: event}
			eventPayload := &connectorpb.EventPayload{Object: alertEventPayload}
			eventPayloads = append(eventPayloads, eventPayload)
		case *connectorpb.Status:
			alertEventPayload := &connectorpb.EventPayload_Status{Status: event}
			eventPayload := &connectorpb.EventPayload{Object: alertEventPayload}
			eventPayloads = append(eventPayloads, eventPayload)
		}
	}

	return eventPayloads
}
