package contract

// Jsonable jsonable
type Jsonable interface {
	ToJSON() ([]byte, error)
}
