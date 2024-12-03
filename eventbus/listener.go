package eventbus

// Listener listener
type Listener interface {
	Handle(event Event) error
}

// EventListener is a listener for a specific event type.
type EventListener[T Event] interface {
	Handle(event T) error
}
