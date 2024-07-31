package queue

import (
	"time"
)

// Job is a job interface
type Job interface {
	// SetQueueable sets the queueable wrapper of the job
	SetQueueable(Queueable)
	// GetQueueable returns the queueable wrapper of the job
	GetQueueable() Queueable
	// GetDelay returns the delay of the job
	GetDelay() time.Duration
	// GetMaxAttempts returns the max count of attempts allowed of the job
	GetMaxAttempts() int
	// GetRetryDelay returns the delay between retries of the job
	GetRetryDelay() time.Duration
	// Handle handles the job
	Handle() error
	// Failed notifies the job that it has failed
	Failed(error)
}
