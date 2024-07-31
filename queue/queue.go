package queue

// Queue is a queue interface
type Queue interface {
	// Name returns the queue name
	Name() string
	// Empty returns true if the queue is empty
	Empty() bool
	// Count returns the number of jobs in the queue
	Count() int64
	// Enqueue adds a job to the queue and returns the job ID
	Enqueue(job Job) (Job, bool)
	// Dequeue dequeues a job from the queue and returns the job
	Dequeue() (Job, bool)
	// Remove removes a job from the queue by ID
	Remove(job Job)
	// Ack acknowledges a job
	Ack(job Job)
	// Release releases a job back to the queue
	Release(job Job)
}
