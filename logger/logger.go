package logger

import (
	"context"
)

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
	// Error logs a message at exception level.
	Error(message string)
	// Errorf logs a formatted message at exception level
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

// LoggerManager is a logger manager and also a logger itself.
type LoggerManager interface {
	Logger
	// SetDefault sets the default channel name.
	SetDefault(name string)
	// HasChannel returns true if the channel with the given name exists.
	HasChannel(name string) bool
	// SetChannel sets the channel with the given name.
	SetChannel(name string, channel Logger)
	// GetChannel returns the channel with the given name.
	GetChannel(name string) Logger
	// GetChannels returns all channels.
	GetChannels() map[string]Logger
	// RemoveChannel removes the channel with the given name.
	RemoveChannel(name string)
	// GetChannelBundle returns a bundle of channels.
	GetChannelBundle(names ...string) Logger
}
