package contract

import "encoding/json"

// List list
type List[E any] interface {
	Countable
	Sortable[E]
	Traversable[int, E]
	Reversible
	Stringable
	Jsonable
	Arrayable[E]
	json.Marshaler
	json.Unmarshaler
	// IsEmpty returns whether the list is empty.
	IsEmpty() bool
	// IsNotEmpty returns whether the list is not empty.
	IsNotEmpty() bool
	// Contains returns whether the list contains the specific element.
	Contains(E) bool
	// ContainsWhere returns whether the list contains specific elements by callback.
	ContainsWhere(func(E) bool) bool
	// Push pushes elements into the list.
	Push(...E)
	// Remove removes the specific element.
	Remove(E)
	// RemoveAt removes the element on the specific index.
	RemoveAt(int)
	// RemoveWhere removes specific elements by callback.
	RemoveWhere(func(E) bool)
	// Clear clears the list.
	Clear()
	// Get returns the element on the specific index.
	Get(int) E
	// Set sets element on the specific index.
	Set(int, E)
	// First returns the first element of the list.
	// it will return a zero value and false when the list is empty.
	First() (E, bool)
	// FirstOr returns the first element of the list, it will return the default value when the list is empty.
	FirstOr(E) E
	// FirstWhere returns the first element of the list which matches the callback.
	// It will return a zero value and false when none matches the callback.
	FirstWhere(func(E) bool) (E, bool)
	// FirstWhereOr returns the first element of the list which matches the callback.
	// It will return the default value when none matches the callback.
	FirstWhereOr(func(E) bool, E) E
	// Last returns the last element of the list.
	// It will return a zero value and false when the list is empty.
	Last() (E, bool)
	// LastOr returns the last element of the list.
	// It will return the default value when the list is empty.
	LastOr(E) E
	// LastWhere returns the last element of the list which matches the callback.
	// It will return a zero value and false when none matches the callback.
	LastWhere(func(E) bool) (E, bool)
	// LastWhereOr returns the last element of the list which matches the callback.
	// It will return the default value when none matches the callback.
	LastWhereOr(func(E) bool, E) E
	// Pop removes the last element of the list and returns it.
	// It will return a zero value and false when the list is empty.
	Pop() (E, bool)
	// Shift removes the first element of the list and returns it.
	// It will return a zero value and false when the list is empty.
	Shift() (E, bool)
	// Unshift puts elements to the head of the list.
	Unshift(...E)
	// IndexOf returns the index of the specific element.
	IndexOf(E) int
	// IndexOfWhere returns the index of the first element which matches the callback.
	IndexOfWhere(func(E) bool) int
	// Compact makes the list more compact
	Compact(eq func(E, E) bool)
	// Min returns the min element
	Min(func(E, E) int) E
	// Max returns the max element
	Max(func(E, E) int) E
}
