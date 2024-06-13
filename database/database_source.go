package database

type DatabaseSource interface {
	DSN() string
}
