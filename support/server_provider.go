package support

import "github.com/gopi-frame/contract/container"

// ServerProvider server provider interface
type ServerProvider interface {
	Register(container.Container)
}
