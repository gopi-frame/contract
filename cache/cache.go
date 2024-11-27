package cache

import (
	"time"
)

// Cache is an interface for caching data.
type Cache interface {
	// Get retrieves the value associated with the given key.
	Get(key string) (string, error)
	// Set sets the value associated with the given key and expiration time.
	Set(key string, value string, expire time.Duration) error
	// Load loads the value associated with the given key.
	// If the value is not found, the loader function is called to load the value and store it in the cache.
	Load(key string, loader func() (value string, err error), expire time.Duration) (string, error)
	// Delete deletes the value associated with the given key.
	Delete(key string) error
	// Has checks if the cache has the given key.
	Has(key string) bool
	// Clear clears the cache.
	Clear() error
}

type CacheManager interface {
	Cache
	SetDefaultStore(name string)
	HasStore(name string) bool
	AddStore(name string, store Cache)
	GetStore(name string) Cache
}
