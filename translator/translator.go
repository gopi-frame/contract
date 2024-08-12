package translator

import (
	"golang.org/x/text/language"
	"io/fs"
	"net/http"
)

// Translator is the interface for creating localizers.
type Translator interface {
	// T translates a message.
	T(id string, data ...any) string
	// P translates a message with a plural count.
	P(id string, pluralCount any, data ...any) string
	// M translates a [Message]
	M(msg Message) string
	// Locale creates a translator for the given languages.
	Locale(languages ...string) Translator
	// AddMessages adds messages to the bundle.
	AddMessages(language string, messages ...Message) error
	// AddMessagesByLanguageTag adds messages to the bundle by language tag.
	AddMessagesByLanguageTag(language language.Tag, messages ...Message) error
	// LoadMessage loads messages from a loader and parser.
	LoadMessage(loader Loader, parser Parser) error
	// LoadMessageFile loads messages from a file.
	LoadMessageFile(path string) error
	// LoadMessageFileFS loads messages from a file from the given file system.
	LoadMessageFileFS(fsys fs.FS, path string) error
	// LoadMessageRemote loads messages from a remote url.
	LoadMessageRemote(remote string, parser Parser) error
	// LoadMessageRemoteRequest loads messages from a custom request
	LoadMessageRemoteRequest(req *http.Request, parser Parser, clientOpts ...func(client *http.Client) error) error
	// RegisterUnmarshalFunc registers a custom unmarshal function for the given format
	RegisterUnmarshalFunc(format string, unmarshaller func(data []byte, v any) error)
	// LanguageTags returns the supported language tags.
	LanguageTags() []language.Tag
}
