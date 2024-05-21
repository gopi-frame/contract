package config

import "io"

// Repository repository interface
type Repository interface {
	Has(module string, key string) bool
	Get(module string, key string, defaultValue ...any) any
	Set(module string, key string, defaultValue any)
	Read(module string, reader io.Reader) error
	Unmarshal(module string, dest any) error
	UnmarshalKey(module string, key string, dest any) error
}
