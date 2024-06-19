package queue

import "time"

type Queue interface {
	Empty() bool
	Count() int64
	Enqueue(JobInterface)
	Dequeue() JobInterface
	Remove(JobInterface)
	Ack(JobInterface)
	Release(JobInterface, time.Duration)
}
