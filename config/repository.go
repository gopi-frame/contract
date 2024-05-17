package config

// Repository repository interface
type Repository interface {
	Has(key string) bool
	Get(key string, defaultValue ...any) any
	Set(key string, defaultValue any)
}
