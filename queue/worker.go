package queue

type Worker interface {
	Handle(job Job)
}
