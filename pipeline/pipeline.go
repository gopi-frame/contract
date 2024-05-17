package pipeline

// Pipeline pipeline interface
type Pipeline[P, R any] interface {
	Send(P) Pipeline[P, R]
	Through(stops ...Stop[P, R]) Pipeline[P, R]
	Then(destination func(P) R) R
}
