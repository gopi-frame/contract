package logger

type Hook interface {
	Enable(level string) bool
	Handle(Entry) error
}
