package eventbus

// Bus interface
type Bus interface {
	// Listen listen
	Listen(topic []string, listeners ...Listener)
	// HasListener has listener
	HasListener(topic string) bool
	// Subscribe adds an subscriber
	Subscribe(subscriber Subscriber)
	// Dispatch dispatches an event
	Dispatch(event Event) error
}
