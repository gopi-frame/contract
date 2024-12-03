package server

import (
	"context"
)

type Driver interface {
	Open(options map[string]any) (Server, error)
}

type Server interface {
	Serve() error
	Shutdown(ctx context.Context) error
	Close() error
}

type RebootableServer interface {
	Server
	Reboot(ctx context.Context) error
}
