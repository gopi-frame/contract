package support

import "encoding/json"

type Set[E Comparable] interface {
	Countable
	Traversable[int, E]
	Arrayable[E]
	Jsonable
	json.Marshaler
	json.Unmarshaler
	Stringable
	IsEmpty() bool
	IsNotEmpty() bool
	Contains(E) bool
	ContainsWhere(func(E) bool) bool
	Push(...E)
	Remove(E)
	RemoveWhere(func(E) bool)
	Clear()
}
