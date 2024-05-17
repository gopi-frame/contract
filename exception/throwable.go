package exception

// Throwable throwable error interface
type Throwable interface {
	error
	SetMessage(message string) Throwable
	StackTrace() string
	SetStackTrace(stackTrace string) Throwable
	Previous() Throwable
	SetPrevious(e Throwable) Throwable
}
