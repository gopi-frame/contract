// Package contract provides interfaces and types for Gopi frame.
package contract

// Arrayable arrayable
type Arrayable[E any] interface {
	// ToArray converts to array
	ToArray() []E
}
