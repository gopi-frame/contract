package queue

import "github.com/google/uuid"

type Queueable interface {
	GetID() uint64
	GetUUID() uuid.UUID
	GetQueue() string
	GetPayload() Job
	GetAttempts() uint
	Fire()
	Fail(error)
}
