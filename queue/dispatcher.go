package queue

type Dispatcher interface {
	FailedJobProvider(FailedJobProvider)
	Dispatch(Job) bool
	Exec()
	Reload()
	Flush()
}
