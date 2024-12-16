package filesystem

type FileSystem interface {
	FileSystemReader
	FileSystemWriter
}

type Driver interface {
	Open(options map[string]any) (FileSystem, error)
}
