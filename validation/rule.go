package validation

type Rule interface {
	Validate(Form) bool
}
