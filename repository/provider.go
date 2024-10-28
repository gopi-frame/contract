package repository

type Provider interface {
	Read() ([]byte, error)
}
