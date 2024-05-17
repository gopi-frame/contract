package codec

// Marshaler marshaler
type Marshaler interface {
	Marshal(any) ([]byte, error)
}
