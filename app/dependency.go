package app

type Dependency interface {
	Name() string
	Description() string
	Depends() []string
}
