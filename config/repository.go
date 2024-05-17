package config

// Repository repository interface
type Repository interface {
	Has(module string, key string) bool
	Get(module string, key string, defaultValue ...any) any
	Set(module string, key string, defaultValue any)
}
