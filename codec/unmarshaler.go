package codec

// Unmarshaler unmarshaler
type Unmarshaler interface {
	Unmarshal([]byte, any) error
}
