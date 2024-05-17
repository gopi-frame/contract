package queue

import "time"

// Job job interface
type Job interface {
	// SetJob set queueable
	SetJob(job Queueable)
	// GetJob get queueable
	GetJob() Queueable
	// Delay delay
	Delay() time.Duration
	// Timeout timeout
	Timeout() time.Duration
	// MaxAttempts max attempts
	MaxAttempts() uint
	// RetryDelay retry delay
	RetryDelay() time.Duration
	// MaxRetryDelay max retry delay
	MaxRetryDelay() time.Duration
	// RetryDelayStep retry delay step
	RetryDelayStep() time.Duration
	// Handle handle
	Handle() error
	// Failed failed
	Failed(err error)
}
