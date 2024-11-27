package logger

type Level interface {
	Enable(level Level) bool
}
