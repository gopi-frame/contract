package logger

import "context"

// Logger logger interface
type Logger interface {
	// WithLevel returns a new logger with level.
	WithLevel(level Level) Logger
	// WithContext returns a new logger with context.
	WithContext(ctx context.Context) Logger
	// Debug logs a message at debug level.
	Debug(message string)
	// Debugf logs a formatted message at debug level.
	Debugf(format string, args ...any)
	// Info logs a message at info level.
	Info(message string)
	// Infof logs a formatted message at info level.
	Infof(format string, args ...any)
	// Warn logs a message at warn level.
	Warn(message string)
	// Warnf logs a formatted message at warn level.
	Warnf(format string, args ...any)
	// Error logs a message at error level.
	Error(message string)
	// Errorf logs a formatted message at error level
	Errorf(format string, args ...any)
	// Panic logs a message at panic level
	Panic(message string)
	// Panicf logs a formatted message at panic level
	Panicf(format string, args ...any)
	// Fatal logs a message at fatal level
	Fatal(message string)
	// Fatalf logs a formatted message at fatal level
	Fatalf(format string, args ...any)
}
