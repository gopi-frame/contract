package translator

type Loader interface {
	Load() ([]byte, error)
}
