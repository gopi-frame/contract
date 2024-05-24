package foundation

import (
	"github.com/gopi-frame/contract/container"
)

// Application application interface
type Application interface {
	container.Container
	BasePath(path string) string
	ConfigPath(path string) string
	StoragePath(path string) string
	Register(service ServiceProvider)
}
