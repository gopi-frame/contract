package queue

import (
	"github.com/gopi-frame/contract/support"
)

// FailedJobProvider failed job provider
type FailedJobProvider interface {
	Save(queue string, job Job, err error)
	All() support.List[Job]
	Find(uuid string) Job
	Forget(uuid string)
	Flush()
}
