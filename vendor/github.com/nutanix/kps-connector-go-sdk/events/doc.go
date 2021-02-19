/*
Package events provides a registry interface that allows for creating, registering, and publishing events
such as status updates and alerts. Events are scraped from the connector on the service domain at a regular
interval. The registry provides the implementation of the `GetEvents` method that gets called each time
the service gets scraped for events.

A registry can be instantiated with the NewRegistry function e.g.
	registry := NewRegistry()

Registering an alert requires instantiating an alert with `NewAlert` and registering it with the registry e.g.
	unableToFetchDataAlert := NewAlert("unableToFetchData", "unable to fetch the data required to stream", connector.Severity_SEVERITY_CRITICAL, connector.State_STATE_FAILED)
	registry.RegisterAlert(unableToFetchDataAlert)

Raising the alert requires calling the `Publish` method on the alert e.g.
	unableToFetchDataAlert.Publish()
	unableToFetchDataAlert.Publish(AlertWithStreamID("..."))
	unableToFetchDataAlert.Publish(AlertWithStreamID("..."), AlertWithMetadata(metadata))

Registering a status requires instantiating an alert with `NewStatus` and registering it with the registry e.g.
	unableToContactDB := NewStatus("unableToContactDB", "unable to contact the database required to stream data", connector.State_STATE_UNHEALTHY)
	registry.RegisterStatus(unableToContactDB)

Updating the status requires calling the `Publish` method on the status e.g.
	unableToContactDB.Publish()
	unableToContactDB.Publish(StatusWithStreamID("..."))
	unableToContactDB.Publish(StatusWithStreamID("..."), AlertWithMetadata(metadata))

Using the registry in the grpc server implementing the contract is as simple as embedding the registry object. e.g.
This makes sure that your server fulfils the `GetEvents` method needed for periodically scraping the events.
	type connector struct {
		*Registry
	}
*/
package events
