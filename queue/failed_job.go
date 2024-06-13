package queue

import "github.com/google/uuid"

type FailedJob interface {
	GetID() uint64
	GetUUID() uuid.UUID
	GetQueue() string
	GetPayload() Job
	GetException() string
}
