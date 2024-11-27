package enum

type Enum interface {
	String() string
	Parse(s string) (Enum, error)
	Equals(other Enum) bool
	Values() []Enum
	Contains(other Enum) bool
}
