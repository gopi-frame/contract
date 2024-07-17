package contract

import "encoding/json"

// Map map
type Map[K comparable, V any] interface {
	Countable
	Traversable[K, V]
	Jsonable
	json.Marshaler
	json.Unmarshaler
	Mappable[K, V]
	Stringable
	// IsEmpty returns whether the map is empty.
	IsEmpty() bool
	// IsNotEmpty returns whether the map is not empty.
	IsNotEmpty() bool
	// Get returns the value of the specific key.
	// It will return a zero value and false when the map is empty.
	Get(K) (V, bool)
	// GetOr returns the value of the specific key.
	// It will return the default value when the map is empty.
	GetOr(K, V) V
	// Set sets value to specific key.
	Set(K, V)
	// Remove removes specific key.
	Remove(K)
	// Keys returns all keys.
	Keys() []K
	// Values returns all values.
	Values() []V
	// Clear clears map.
	Clear()
	// ContainsKey returns whether the map contains specific key.
	ContainsKey(K) bool
	// Contains returns whether the map contains specific value.
	Contains(V) bool
	// ContainsWhere returns whether the map contains value which matches the callback.
	ContainsWhere(func(V) bool) bool
}
