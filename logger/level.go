package logger

type Level int8

const (
	LevelDebug = iota - 1
	LevelInfo
	LevelWarn
	LevelError
	LevelPanic
	LevelFatal
)
