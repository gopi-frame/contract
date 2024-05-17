package logger

import (
	"time"
)

// Entry entry
type Entry interface {
	Timestamp() time.Time
	File() string
	Line() int
	Level() string
	Message() string
	Fields() map[string]any
}
