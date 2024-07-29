package queue

// Queueable is a queueable interface, it is used to wrap a job
type Queueable interface {
	// GetID returns the job ID
	GetID() string
	// GetQueue returns the job queue
	GetQueue() string
	// GetPayload returns the job payload
	GetPayload() Job
	// GetAttempts returns how many times the job has been attempted
	GetAttempts() int
}
