package pipeline

// Stop stop interface
type Stop[P, R any] interface {
	Handle(P, func(P) R) R
}
