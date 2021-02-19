package events

import (
	"github.com/golang/glog"
	"google.golang.org/protobuf/types/known/structpb"
)

// EventMetadata is a mechanism for adding extra arbitrary properties to an event as metadata
type EventMetadata struct {
	ErrorMessage string
	StreamID     string
	Extra        map[string]interface{}
}

const (
	errorMessageProp = "ErrorMessage"
	streamIDProp     = "StreamID"
	extraMessageProp = "ExtraMessage"
)

func (em *EventMetadata) toStruct() (*structpb.Struct, error) {
	m := make(map[string]interface{})
	m[errorMessageProp] = em.ErrorMessage
	m[streamIDProp] = em.StreamID
	m[extraMessageProp] = em.Extra

	metadata, err := structpb.NewStruct(m)
	if err != nil {
		glog.Errorf("unable convert properties to metadata proto structure")
		return nil, err
	}

	return metadata, nil
}
