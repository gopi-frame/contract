package config

// Repository repository interface
type Repository interface {
	All() map[string]any
	Has(path string) bool
	Get(path string, defalutValue ...any) any
	Set(path string, value any)
	Unmarshal(dest any) error
	UnmarshalKey(key string, dest any) error
}
