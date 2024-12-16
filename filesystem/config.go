package filesystem

type CreateDirectoryConfig interface {
	DirVisibility() string
}

type CreateFileConfig interface {
	CreateDirectoryConfig
	WriteFlag() int
	FileVisibility() string
}
