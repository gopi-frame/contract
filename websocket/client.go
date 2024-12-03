package websocket

import "github.com/gorilla/websocket"

type Client interface {
	Conn() *websocket.Conn
	SendMessage(messageType int, data []byte) error
	SendJSON(data any) error
	Close() error
	Closed() bool
}

type ClientManager interface {
	Add(client Client) string
	Get(id string) (Client, error)
	Remove(iid string) error
	RemoveAll() error
	Range(f func(key string, value Client) bool)
}
