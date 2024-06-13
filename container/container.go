package container

// Container application interface
type Container interface {
	// Set register an existing instance
	Set(abstract string, instance any)
	// Has has
	Has(string) bool
	// Get get an instance
	Get(string) any
	// Bind bind abstract an concrate
	Bind(abstract string, concrate func() any)
	// BindIf bind if abstract is not bound
	BindIf(abstract string, concrate func() any)
	// Alias set alias for name
	Alias(abstract string, alias string)
	// Extend extend
	Extend(abstract string, extender ...func(instance any) any)
}
