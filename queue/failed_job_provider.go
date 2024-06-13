package queue

import "github.com/google/uuid"

type FailedJobProvider interface {
	Find(uuid uuid.UUID) FailedJob
	All(queue string) []FailedJob
	Forget(uuid uuid.UUID)
	Flush(queue string)
}
