package contract

import "time"

// Delayable delayable
type Delayable[T any] interface {
	Until() time.Time
	Value() T
}
