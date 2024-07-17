package contract

// Arrayable arrayable
type Arrayable[E any] interface {
	// ToArray converts to array
	ToArray() []E
}
