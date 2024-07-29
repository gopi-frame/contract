package queue

import "time"

// Queue is a queue interface
type Queue interface {
	Name() string
	Empty() bool
	Count() int64
	Enqueue(Job) bool
	Dequeue() Job
	Remove(Job)
	Ack(Job)
	Release(Job, time.Duration)
}
