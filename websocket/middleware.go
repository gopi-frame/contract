package websocket

type Middleware interface {
	Handle(client Client, message Message, next Handler)
}
