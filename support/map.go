package support

import "encoding/json"

// Map map interface
type Map[K comparable, V any] interface {
	Countable
	Traversable[K, V]
	Jsonable
	json.Marshaler
	json.Unmarshaler
	Mapable[K, V]
	Stringable
	IsEmpty() bool
	IsNotEmpty() bool
	Get(K) (V, bool)
	GetOr(K, V) V
	Set(K, V)
	Remove(K)
	Keys() []K
	Values() []V
	Clear()
	ContainsKey(K) bool
	Contains(V) bool
	ContainsWhere(func(V) bool) bool
}
