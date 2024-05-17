package event

// Subscriber subscriber
type Subscriber interface {
	// Subscribe subscribe
	Subscribe(dispatcher Dispatcher)
}
