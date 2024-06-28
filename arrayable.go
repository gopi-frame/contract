package contract

// Arrayable arrayable
type Arrayable[E any] interface {
	ToArray() []E
}
