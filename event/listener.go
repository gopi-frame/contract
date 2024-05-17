package event

// Listener listener
type Listener interface {
	Handle(event Event) bool
}
