package validation

import "github.com/gopi-frame/contract/support"

// Form form interface
type Form interface {
	SetLocale(locale string)
	Locale() string
	AddError(key string, message string)
	Fails() bool
	Errors() map[string][]string
	CustomRules() support.List[Rule]
}
