package eventbus

// Listener listener
type Listener interface {
	Handle(event Event) error
}
