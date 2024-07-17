package contract

// Jsonable jsonable
type Jsonable interface {
	// ToJSON converts to json
	ToJSON() ([]byte, error)
}
