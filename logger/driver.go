package logger

// Driver logger driver
type Driver interface {
	Open(map[string]any) (Logger, error)
}
