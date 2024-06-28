package contract

import "encoding/json"

// List list interface
type List[E any] interface {
	Countable
	Sortable[E]
	Traversable[int, E]
	Reversable
	Stringable
	Jsonable
	Arrayable[E]
	json.Marshaler
	json.Unmarshaler
	IsEmpty() bool
	IsNotEmpty() bool
	Contains(E) bool
	ContainsWhere(func(E) bool) bool
	Push(...E)
	Remove(E)
	RemoveAt(int)
	RemoveWhere(func(E) bool)
	Clear()
	Get(int) E
	Set(int, E)
	First() (E, bool)
	FirstOr(E) E
	FirstWhere(func(E) bool) (E, bool)
	FirstWhereOr(func(E) bool, E) E
	Last() (E, bool)
	LastOr(E) E
	LastWhere(func(E) bool) (E, bool)
	LastWhereOr(func(E) bool, E) E
	Pop() (E, bool)
	Shift() (E, bool)
	Unshift(...E)
	IndexOf(E) int
	IndexOfWhere(func(E) bool) int
	Compact(eq func(E, E) bool)
	Min(func(E, E) int) E
	Max(func(E, E) int) E
}
