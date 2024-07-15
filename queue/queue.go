package queue

import "time"

type Queue interface {
	Name() string
	Empty() bool
	Count() int64
	Enqueue(Job)
	Dequeue() Job
	Remove(Job)
	Ack(Job)
	Release(Job, time.Duration)
}
