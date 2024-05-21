package support

import "github.com/gopi-frame/contract/container"

// ServerProvider server provider interface
type ServerProvider interface {
	// Register register
	Register(container.Container)
	// Boot boot
	Boot(container.Container)
}
