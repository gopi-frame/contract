package repository

type Parser interface {
	Unmarshal(data []byte) (map[string]any, error)
}
