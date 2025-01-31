package app

// Component is the interface for gopi app component.
type Component interface {
	Name() string
	Register(App) error
	Unregister(App) error
	Boot() error
	Booted() bool
	Shutdown() error
}
