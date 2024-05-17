package validation

// Engine validation engine interface
type Engine interface {
	ValidateForm(form Form)
}
