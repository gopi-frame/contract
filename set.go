package contract

import "encoding/json"

// Set sets
type Set[E comparable] interface {
	Countable
	Traversable[int, E]
	Arrayable[E]
	Jsonable
	json.Marshaler
	json.Unmarshaler
	Stringable
	// IsEmpty returns whether the set is empty
	IsEmpty() bool
	// IsNotEmpty returns whether the set is not empty
	IsNotEmpty() bool
	// Contains returns whether the set contains specific element
	Contains(E) bool
	// ContainsWhere returns whether the set contains element which matches the callback
	ContainsWhere(func(E) bool) bool
	// Push pushes elements to set
	Push(...E)
	// Remove removes the specific element
	Remove(E)
	// RemoveWhere removes elements which matches the callback
	RemoveWhere(func(E) bool)
	// Clear clears the set
	Clear()
}
