package support

// Jsonable jsonable
type Jsonable interface {
	ToJSON() ([]byte, error)
}
