package redis

// Driver is a redis driver interface
type Driver interface {
	Open(options map[string]any) (Client, error)
}
