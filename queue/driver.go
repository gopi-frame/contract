package queue

type Driver interface {
	Open(options map[string]any) (Queue, error)
}
