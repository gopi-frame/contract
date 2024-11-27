package console

type Value interface {
	String() string
	Set(string) error
	Type() string
}

type SliceValue interface {
	Append(value string) error
	Replace(values []string) error
	GetSlice() []string
}

type Valuer[T any] interface {
	Value() T
}
