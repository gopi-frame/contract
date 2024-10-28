package app

import "github.com/spf13/cobra"

// Component is the interface for gopi app component.
type Component interface {
	Name() string
	Register(app App)
	Depends() []string
	Boot()
	Booted() bool
}

// PublishableComponent is the interface for gopi app component that can publish files.
type PublishableComponent interface {
	Component
	Publishes() map[string]string
}

// ExecutableComponent is the interface for gopi app component that can execute commands.
type ExecutableComponent interface {
	Component
	Commands() []*cobra.Command
}
