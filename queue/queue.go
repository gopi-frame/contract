package queue

import (
	"github.com/gopi-frame/contract/event"
)

// Queue queue
type Queue interface {
	// Name return queue name
	Name() string
	// Count returns the count of pending jobs
	Count() int64
	// IsEmpty returns if the count of pending jobs is zero
	IsEmpty() bool
	// Enqueue pushes a job to queue
	Enqueue(job Job) bool
	// Dequeue pops a job from queue
	Dequeue() (Job, bool)
	// Remove removes a job from queue
	Remove(job Job) bool
	// Ack acks a job
	Ack(job Job)
	// Fail handles a failed job
	Fail(job Job, err error)
	// Flush removes all failed jobs
	Flush()
	// Reload reloads all failed jobs into queue
	Reload()
	// Subscribe add a subscriber to queue events
	Subscribe(subscriber Subcriber)
	// DispatchEvent dispatches specifia event
	DispatchEvent(event event.Event)
	// Progress progress
	Progress() (pending int64, executing int64)
}

// AbstractQueue abstract queue
type AbstractQueue struct {
	Queue
	events event.Dispatcher
}

// Subscribe subscribe
func (q *AbstractQueue) Subscribe(subscriber event.Subscriber) {
	if q.events != nil {
		q.events.Subscribe(subscriber)
	}
}

// DispatchEvent dispatch event
func (q *AbstractQueue) DispatchEvent(event event.Event) {
	if q.events != nil {
		q.events.Dispatch(event)
	}
}
