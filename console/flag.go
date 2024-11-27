package console

type Flag interface {
	Name() string
	Shorthand() string
	Usage() string
	Type() string
	Value() Value
	Hidden() bool
	IsBool() bool
}
