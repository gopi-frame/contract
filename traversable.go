package contract

// Traversable traversable
type Traversable[K any, V any] interface {
	// Each travers, if the callback returns false then break
	Each(func(key K, value V) bool)
}
