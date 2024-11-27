package validation

// Param is a interface for error parameters.
type Param interface {
	Key() string
	Value() string
}

// ErrorBuilder is a interface for building messages.
type ErrorBuilder interface {
	BuildError(code string, message string, params ...Param) Error
}

// Error is a interface for errors.
type Error interface {
	error
	Code() string
	Message() string
	SetMessage(message string) Error
	Params() []Param
	HasParam(key string) bool
	AddParam(param Param) Error
}

type Errors interface {
	Each(f func(code string, err Error) bool)
	Has(code string) bool
	Add(err Error)
	Get(code string) Error
}

// ErrorBag is a interface for error bags.
type ErrorBag interface {
	Error
	Each(f func(path string, errs Errors) bool)
	// Fails returns true if the error bag has one or more errors.
	Fails() bool
	// Failed returns the keys of the failed fields.
	Failed() []string
	// FailedAt returns true if the error bag has one or more errors for the given keys.
	FailedAt(key string, codes ...string) bool
	// HasError returns true if the error bag has one or more errors for the given keys.
	HasError(key string) bool
	// AddError adds an error to the error bag.
	AddError(key string, err Error)
	// GetAllErrors returns the errors in the error bag.
	GetAllErrors() map[string]Errors
	// GetErrors returns the errors for the given key.
	GetErrors(key string) Errors
	// GetError returns the error for the given key and code.
	GetError(key string, code string) Error
	// SetMessages sets custom error messages
	SetMessages(messages map[string]map[string]string) ErrorBag
	// GetMessages returns the messages in the error bag.
	GetMessages() map[string][]string
	// GetMessage returns the messages for the given key.
	GetMessage(key string) []string
}
