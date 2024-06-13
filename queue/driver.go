package queue

type Queue interface {
	Queue() string
	Count() int64
	Empty() bool
	Enqueue(Job)
	Dequeue() Job
	Remove(Job)
}
