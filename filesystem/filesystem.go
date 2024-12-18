package filesystem

type FileSystem interface {
	FileSystemReader
	FileSystemWriter
}

type FileSystemManager interface {
	FileSystem
	AddFS(name string, fs FileSystem)
	TryGetFS(name string) (FileSystem, error)
	GetFS(name string) FileSystem
	HasFS(name string) bool
}

type Driver interface {
	Open(options map[string]any) (FileSystem, error)
}
