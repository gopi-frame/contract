package translator

import "golang.org/x/text/language"

// Message is the interface of message which can be localized.
type Message interface {
	// GetID returns the unique identity of the message.
	GetID() string
	// GetHash returns the unique identity of the message content.
	GetHash() string
	// GetDescription returns the description of the message.
	GetDescription() string
	// GetLeftDelim returns the left Go template delimiter.
	GetLeftDelim() string
	// GetRightDelim returns the right Go template delimiter.
	GetRightDelim() string
	// GetZero returns the content of the message for the CLDR plural form "zero".
	GetZero() string
	// GetOne returns the content of the message for the CLDR plural form "one".
	GetOne() string
	// GetTwo returns the content of the message for the CLDR plural form "two".
	GetTwo() string
	// GetFew returns the content of the message for the CLDR plural form "few".
	GetFew() string
	// GetMany returns the content of the message for the CLDR plural form "many".
	GetMany() string
	// GetOther returns the content of the message for the CLDR plural form "other".
	GetOther() string
}

// MessagePack is the interface of message pack which can be localized.
type MessagePack interface {
	GetMessages() []Message
	GetLanguageTag() language.Tag
}
