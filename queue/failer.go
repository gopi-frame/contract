package queue

import (
	"github.com/google/uuid"
	"github.com/gopi-frame/contract/support"
)

// Failer failed job provider
type Failer interface {
	Save(queue string, job Job, err error)
	All(queue string) support.List[Job]
	Find(queue string, id uuid.UUID) Job
	Forget(queue string, id uuid.UUID)
	Flush(queue string)
}
