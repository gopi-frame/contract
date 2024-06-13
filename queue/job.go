package queue

import "time"

type Job interface {
	SetQueueable(Queueable)
	Queueable() Queueable
	MaxAttempts() uint
	AvalidableAt() time.Time
	Handle() error
	Failed(err error)
}
