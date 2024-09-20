package validation

import "context"

// Validator is the validator interface.
type Validator interface {
	// Validate validates the given value.
	Validate(ctx context.Context, builders ...ValidatorBuilder) ErrorBag
}

// ValidatorContext is the context for validation.
type ValidatorContext interface {
	AddValidate(key string, rule Validatable)
}

// ValidatorBuilder is the validator builder interface.
type ValidatorBuilder interface {
	// Build builds the validator.
	Build(ctx ValidatorContext)
	// SetAttribute sets the attribute name.
	SetAttribute(attribute string) ValidatorBuilder
	// GetAttribute gets the attribute name.
	GetAttribute() string
	// SetKey sets the attribute key parts.
	SetKey(paths ...string) ValidatorBuilder
	// GetKey gets the attribute key parts.
	GetKey() []string
}

// Validatable is the interface for creating validatable types on the client side.
type Validatable interface {
	Validate(ctx context.Context, errorBuilder ErrorBuilder) Error
}

type Rule[T any] interface {
	Validate(ctx context.Context, errorBuilder ErrorBuilder, value T) Error
}
