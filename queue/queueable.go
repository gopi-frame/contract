package queue

type Queueable interface {
	GetID() string
	GetQueue() string
	GetPayload() JobInterface
	GetAttempts() int
}
