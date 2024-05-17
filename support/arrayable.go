package support

// Arrayable arrayable
type Arrayable[E any] interface {
	ToArray() []E
}
