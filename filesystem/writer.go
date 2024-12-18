package filesystem

import (
	"io"
)

type FileSystemWriter interface {
	Write(path string, content []byte, config map[string]any) error
	WriteStream(path string, stream io.Reader, config map[string]any) error
	SetVisibility(path string, visibility string) error
	Delete(path string) error
	DeleteDir(path string) error
	CreateDir(path string, config map[string]any) error
	Move(src string, dst string, config map[string]any) error
	Copy(src string, dst string, config map[string]any) error
}
