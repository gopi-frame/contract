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
	IsEmpty() bool
	IsNotEmpty() bool
	Clear()
	Peek() (E, bool)
	Enqueue(E) bool
	Dequeue() (E, bool)
	Remove(E)
	RemoveWhere(func(E) bool)
}

// BlockingQueue blocking queue interface
type BlockingQueue[E any] interface {
	Queue[E]
	TryEnqueue(E) bool
	TryDequeue() (E, bool)
	EnqueueTimeout(E, time.Duration) bool
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
