package logger

// Formatter formatter
type Formatter interface {
	Format(Entry) string
}
