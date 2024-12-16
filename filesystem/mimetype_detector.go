package filesystem

import "io"

type MimeTypeDetector interface {
	Detect(location string, content []byte) string
	DetectFromFile(location string) string
	DetectFromPath(location string) string
	DetectFromReader(reader io.Reader) string
	DetectFromContent(content []byte) string
}
