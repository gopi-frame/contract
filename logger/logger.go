package logger

import (
	"io"

	"github.com/gopi-frame/contract/event"
)

// Logger logger
type Logger interface {
	// Dispatcher set event dispatcher
	Dispatcher(event.Dispatcher)
	// Formatter set formatter
	Formatter(Formatter)
	// Hook set hook
	Hooks(...Hook)
	// Output set outputs
	Outputs(...io.Writer)
	// Debug debug
	Debug(message string, fields map[string]any)
	// Info info
	Info(message string, fields map[string]any)
	// Warn warn
	Warn(message string, fields map[string]any)
	// Error error
	Error(message string, fields map[string]any)
	// Fatal fatal
	Fatal(message string, fields map[string]any)
	// Panic panic
	Panic(message string, fields map[string]any)
}
