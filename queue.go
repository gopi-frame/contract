package contract

import (
	"encoding/json"
	"time"
)

// Queue queue interface
type Queue[E any] interface {
	Countable
	Arrayable[E]
	Jsonable
	json.Marshaler
	json.Unmarshaler
	Stringable
	// IsEmpty returns whether the queue is empty
	IsEmpty() bool
	// IsNotEmpty returns whether the queue is not empty
	IsNotEmpty() bool
	// Clear clears the queue
	Clear()
	// Peek returns the first element of the queue.
	Peek() (E, bool)
	// Enqueue enqueues element to the end of the queue
	Enqueue(E) bool
	// Dequeue dequeues the first element
	Dequeue() (E, bool)
	// Remove removes the specific element
	Remove(E)
	// RemoveWhere removes the elements which matches the callback
	RemoveWhere(func(E) bool)
}

// BlockingQueue blocking queue interface
type BlockingQueue[E any] interface {
	Queue[E]
	// TryEnqueue tries to enqueue element to the end of the queue
	TryEnqueue(E) bool
	// TryDequeue tries to dequeue the first element and return it
	TryDequeue() (E, bool)
	// EnqueueTimeout enqueues element to the end of the queue,
	// it will return false when timeout
	EnqueueTimeout(E, time.Duration) bool
	// DequeueTimeout dequeues the first element,
	// it will return zero value and false when timeout
	DequeueTimeout(time.Duration) (E, bool)
}

type DelayedQueue[E any] interface {
	Countable
	Arrayable[Delayable[E]]
	Jsonable
	json.Marshaler
	json.Unmarshaler
	Stringable
	IsEmpty() bool
	IsNotEmpty() bool
	Clear()
	Peek() (Delayable[E], bool)
	Enqueue(Delayable[E]) bool
	Dequeue() (Delayable[E], bool)
	TryDequeue() (Delayable[E], bool)
	DequeueTimeout(time.Duration) (Delayable[E], bool)
}
