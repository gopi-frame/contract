package eventbus

// Bus interface
type Bus interface {
	// Listen registers a list of listeners to given topics
	Listen(topic []string, listeners ...Listener)
	// HasListener checks if any listeners have been registered to the specific topic
	HasListener(topic string) bool
	// Subscribe registers a subscriber
	Subscribe(subscriber Subscriber)
	// Dispatch dispatches an event
	Dispatch(event Event) error
}
