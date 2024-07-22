package router

type Route interface {
	Name(string) Route
	Use(...Middleware) Route
}
