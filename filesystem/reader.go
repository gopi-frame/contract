package filesystem

import (
	"io"
	"io/fs"
	"os"
	"time"
)

type FileSystemReader interface {
	Exists(path string) (bool, error)
	FileExists(path string) (bool, error)
	DirExists(path string) (bool, error)
	Read(path string) ([]byte, error)
	ReadStream(path string) (io.ReadCloser, error)
	ReadDir(path string) ([]os.DirEntry, error)
	WalkDir(path string, walkFn fs.WalkDirFunc) error
	LastModified(path string) (time.Time, error)
	FileSize(path string) (int64, error)
	MimeType(path string) (string, error)
	Visibility(path string) (string, error)
}
