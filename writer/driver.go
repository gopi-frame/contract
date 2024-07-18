package writer

import "io"

// Driver writer driver interface
type Driver interface {
	Open(options map[string]any) (io.WriteCloser, error)
}
