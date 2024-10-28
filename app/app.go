package app

import (
	"github.com/gopi-frame/contract/container"
	"github.com/gopi-frame/contract/repository"
)

// App is the interface for the application.
type App interface {
	container.Container[any]
	// Name returns the name of the application.
	Name() string
	// Version returns the version of the application.
	Version() string
	// Run runs the application.
	Run() error
	// Debug returns true if the application is in debug mode.
	Debug() bool
	// Root returns the root path of the application.
	Root() string
	// WorkingDirectory returns the working path of the application.
	WorkingDirectory() string
	// StoragePath returns the storage path of the application.
	StoragePath() string
	// ConfigPath returns the config path of the application.
	ConfigPath() string
	// Configure registers a component configuration.
	Configure(component, path string)
	// Config returns the config of the application.
	Config(components ...string) repository.Repository
	// Register registers a provider.
	Register(component Component)
}
