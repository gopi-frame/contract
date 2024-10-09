package cache

// Driver is an interface for cache drivers.
type Driver interface {
	// Open opens the cache driver.
	Open(config map[string]any) (Cache, error)
}
