package connector

import (
	"context"
	"log"

	connectorpb "github.com/nutanix/kps-connector-go-sdk/connector/v1"
	"github.com/nutanix/kps-connector-go-sdk/events"
	"github.com/nutanix/kps-connector-go-sdk/transport"
)

func (d *Connector) getStreams(context.Context) (*connectorpb.GetPayloadResponse, error) {
	d.mtx.RLock()
	defer d.mtx.RUnlock()

	payloads := make([]*connectorpb.Payload, 0)
	for _, stream := range d.streams {
		payloadStream := &connectorpb.Payload_Stream{
			Stream: stream,
		}
		payload := &connectorpb.Payload{Object: payloadStream}
		payloads = append(payloads, payload)
	}
	resp := &connectorpb.GetPayloadResponse{
		Status:   &connectorpb.ResponseStatus{Code: connectorpb.ResponseCode_RESPONSE_CODE_OK},
		Payloads: payloads,
	}

	return resp, nil
}

func (d *Connector) setStreams(ctx context.Context, streams []*connectorpb.Stream) error {
	currStreams := make(map[string]bool)
	log.Printf("number of streams to stream: %d", len(streams))
	for _, stream := range streams {
		currStreams[stream.Id] = true
		switch stream.Direction {
		case connectorpb.StreamDirection_STREAM_DIRECTION_INGRESS:
			err := d.setStreamToTransport(ctx, stream)
			if err != nil {
				return err
			}
		case connectorpb.StreamDirection_STREAM_DIRECTION_EGRESS:
			err := d.setStreamFromTransport(ctx, stream)
			if err != nil {
				return err
			}
		}
	}

	// Unsubscribe from streams that are no longer being used
	for streamID, cancelfunc := range d.activeInStreams {
		inUse := false
		for id := range currStreams {
			if streamID == id {
				inUse = true
			}
		}
		if !inUse {
			log.Printf("cancel context for stream %s", streamID)
			cancelfunc()
			delete(d.activeInStreams, streamID)
		}
	}
	for streamID, subscription := range d.activeOutStreams {
		inUse := false
		for id := range currStreams {
			if streamID == id {
				inUse = true
			}
		}
		if !inUse {
			log.Printf("stopping streaming stream %s", streamID)
			delete(d.activeOutStreams, streamID)
			err := subscription.Unsubscribe()
			if err != nil {
				log.Println(err)
				_ = transportUnsubscribeFailedAlert.Publish(events.AlertWithStreamID(streamID), events.AlertWithEventMetadata(&events.EventMetadata{
					ErrorMessage: err.Error(),
					StreamID:     streamID,
				}))

				return err
			}
		}
	}

	d.streams = streams
	return nil
}

func (d *Connector) setStreamToTransport(ctx context.Context, stream *connectorpb.Stream) error {
	log.Printf("setStreamToTransport: %+v", stream)
	if _, ok := d.activeInStreams[stream.Id]; ok {
		log.Printf("stream already streaming %s", stream.Id)
		return nil
	}

	ctx, cancelfunc := context.WithCancel(context.Background())
	d.activeInStreams[stream.Id] = cancelfunc

	metadata := mapToStreamMetadata(stream.GetMetadata().AsMap())
	consumer := newConsumer()
	if err := consumer.subscribe(ctx, metadata); err != nil {
		_ = streamUnhealthyStatus.Publish(events.StatusWithStreamID(stream.GetId()), events.StatusWithEventMetadata(&events.EventMetadata{
			StreamID:     stream.GetId(),
			ErrorMessage: err.Error(),
		}))
		return err
	}

	tclt, err := transport.NewTransportClient()
	if err != nil {
		_ = streamUnhealthyStatus.Publish(events.StatusWithStreamID(stream.GetId()), events.StatusWithEventMetadata(&events.EventMetadata{
			StreamID:     stream.GetId(),
			ErrorMessage: err.Error(),
		}))
		return err
	}

	go consumerLoop(ctx, stream, consumer, tclt)

	log.Printf("starting streaming stream %s", stream.Id)
	return nil
}

func consumerLoop(ctx context.Context, stream *connectorpb.Stream, c *consumer, tclt transport.Client) {
	_ = streamStartedStatus.Publish(events.StatusWithStreamID(stream.GetId()))
	for {
		select {
		case <-ctx.Done():
			log.Printf("stopping streaming stream %s", stream.GetId())
			return
		default:
			nextMsg, err := c.nextMsg()
			if err != nil {
				_ = streamUnhealthyStatus.Publish(events.StatusWithStreamID(stream.GetId()), events.StatusWithEventMetadata(&events.EventMetadata{
					StreamID:     stream.GetId(),
					ErrorMessage: err.Error(),
				}))
				continue
			}
			msg := transport.Message{
				Payload: nextMsg,
			}
			err = tclt.Publish(stream.GetTransportChannel(), msg)
			if err != nil {
				log.Println(err)
				_ = transportPublishFailedAlert.Publish(events.AlertWithStreamID(stream.GetId()), events.AlertWithEventMetadata(&events.EventMetadata{
					ErrorMessage: err.Error(),
					StreamID:     stream.GetId(),
				}))
				continue
			}
			_ = streamHealthyStatus.Publish(events.StatusWithStreamID(stream.GetId()))
			log.Printf("msg sent: %s", string(msg.Payload))
		}
	}
}

func (d *Connector) setStreamFromTransport(ctx context.Context, stream *connectorpb.Stream) error {
	log.Printf("setStreamFromTransport: %+v", stream)
	log.Printf("starting streaming stream %s", stream.Id)
	if _, ok := d.activeOutStreams[stream.Id]; ok {
		log.Printf("stream already streaming %s", stream.Id)
		return nil
	}

	tclt, err := transport.NewTransportClient()
	if err != nil {
		return err
	}
	streamProducer := newProducer()
	streamMeta := mapToStreamMetadata(stream.Metadata.AsMap())
	if err := streamProducer.connect(ctx, streamMeta); err != nil {
		return err
	}
	sub, err := tclt.Subscribe(stream.GetTransportChannel(), streamProducer.subscribeMsgHandler)
	if err != nil {
		log.Println(err)
		_ = transportSubscribeFailedAlert.Publish(events.AlertWithStreamID(stream.GetId()), events.AlertWithEventMetadata(&events.EventMetadata{
			ErrorMessage: err.Error(),
			StreamID:     stream.GetId(),
		}))
		return err
	}
	d.activeOutStreams[stream.Id] = sub

	return nil
}
