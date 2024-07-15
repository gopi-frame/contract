package eventbus

// Event event interface
type Event interface {
	Topic() string
}
