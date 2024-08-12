package validation

// Rule is the interface for validation rules.
type Rule[T any] interface {
	Validate(value T) error
}
