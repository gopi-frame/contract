package contract

// Mapable mapable
type Mapable[K comparable, V any] interface {
	ToMap() map[K]V
}
