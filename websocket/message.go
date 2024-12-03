package websocket

type Message interface {
	Event() string
	Payload() any
}

type MessageParser interface {
	Parse(messageType int, data []byte) (Message, error)
}
