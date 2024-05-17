package support

// Traversable traversable
type Traversable[K any, V any] interface {
	Each(func(key K, value V) bool)
}
