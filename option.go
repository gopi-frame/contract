package contract

type Option[T any] interface {
	Apply(T) error
}
