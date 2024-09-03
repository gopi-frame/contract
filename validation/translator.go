package validation

// Translator is the interface for translation.
type Translator interface {
	T(key string, params map[string]any) string
	P(key string, pluralCount any, params map[string]any) string
	Locale(locale string) Translator
}
