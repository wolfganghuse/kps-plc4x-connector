/*
Package transport provides a Publish/Subscribe interface for publishing data from streams to data pipelines
or subscribing to data from the data pipelines to publish to the streams.

transport package exposes two interfaces:
	type Client interface {
		Publish(channel string, msg Message) error
		Subscribe(channel string, callback MessageHandler) (Subscription, error)
	}
and
	type Subscription interface {
		Unsubscribe() error
		Channel() string
	}

A `Client` can be created by calling the `NewTransportClient` function:
	client, err := NewTransportClient()

Note, the client created by the `NewTransportClient` function is a singleton. Repeated calls to the function
will return the same client.

In order to publish data into the transport, we need to create a Message object:
	msg := &Message{
		Payload: []byte("example")
	}

Once, the client has been created, it can be used to Publish data from streams into the transport channel:
	client.Publish(stream.GetTransportChannel(), msg)

In order to create a subscription, the client needs to provide a callback with the MessageHandler signature.
This callback gets called with a message as parameter each time a new message is received on the subscribed channel.
	func msgHandler (msg *Message) {
		// Do stuff
	}

Subsequently, a subscription can be created for subscribing to the data from the transport channel:
	sub, err := client.Subscribe(stream.GetTransportChannel(), msgHandler)

A subscription also exposes the channel it is subscribed to via the `Channel` method:
	channel := sub.Channel()

When a subscription is no longer needed, it can be unsubscribed by calling the `Unsubscribe` method:
	sub.Unsubscribe()
*/
package transport
