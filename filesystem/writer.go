package filesystem

import (
	"io"
)

type FileSystemWriter interface {
	Write(path string, content []byte, config CreateFileConfig) error
	WriteStream(path string, stream io.Reader, config CreateFileConfig) error
	SetVisibility(path string, visibility string) error
	Delete(path string) error
	DeleteDir(path string) error
	CreateDir(path string, config CreateDirectoryConfig) error
	Move(src string, dst string, config CreateDirectoryConfig) error
	Copy(src string, dst string, config CreateFileConfig) error
}
