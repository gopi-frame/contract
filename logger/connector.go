package logger

type Connector interface {
	Connect(map[string]any) Logger
}