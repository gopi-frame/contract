package contract

// Mappable mappable
type Mappable[K comparable, V any] interface {
	// ToMap converts to map
	ToMap() map[K]V
}
