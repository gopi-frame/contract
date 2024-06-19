package queue

import "time"

type JobInterface interface {
	SetModel(Queueable)
	GetModel() Queueable
	GetDelay() time.Duration
	GetMaxAttempts() int
	GetRetryDelay() time.Duration
	Handle() error
	Failed(error)
}

type Job struct {
	JobInterface
	model Queueable
}

func (j *Job) SetModel(model Queueable) {
	j.model = model
}

func (j *Job) GetModel() Queueable {
	return j.model
}

func (j *Job) Release(delay time.Duration) {
}
