package queue

import "github.com/google/uuid"

// Queueable queueable
type Queueable interface {
	ID() uint64
	UUID() uuid.UUID
	Attempts() uint
	Payload() Job
}
