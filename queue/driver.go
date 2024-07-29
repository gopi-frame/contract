package queue

// Driver is a queue driver interface
type Driver interface {
	Open(options map[string]any) (Queue, error)
}
