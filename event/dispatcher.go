package event

// Dispatcher interface
type Dispatcher interface {
	// Listen listen
	Listen(topic []Event, listeners ...Listener)
	// HasListener has listener
	HasListener(topic Event) bool
	// Subscribe adds an subscriber
	Subscribe(subscriber Subscriber)
	// Dispatch dispatches an event
	Dispatch(event Event)
}
