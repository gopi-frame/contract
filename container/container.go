package container

// Container container
type Container[T any] interface {
	// Set sets a value in the container by name
	Set(name string, value T)
	// Get gets a value from the container by name or creates a new one firstly if not found.
	// If you want to always get a new instance of the value, use [Make] instead.
	Get(name string) T
	// GetE gets a value from the container by name but returns exception if the value is not found or cannot be created.
	// If you want to always get a new instance of the value, use [Make] instead.
	GetE(name string) (T, error)
	// Has checks if the container has a value or a constructor by name
	Has(name string) bool
	// Lazy sets a constructor function that will be called when the value is accessed
	Lazy(name string, constructor func() (T, error))
	// Make makes a new instance of the value, If the value is a constructor function, it will call it
	Make(name string) (T, error)
}
