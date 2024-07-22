package router

import (
	"net/http"

	"github.com/gopi-frame/contract/response"
)

type Handler = func(request *http.Request) response.Responser

type Router interface {
	http.Handler
	Use(...Middleware) Router
	Group(RouteGroup, func(Router)) Router
	GET(path string, handler Handler) Route
	POST(path string, handler Handler) Route
	PUT(path string, handler Handler) Route
	PATCH(path string, handler Handler) Route
	DELETE(path string, handler Handler) Route
	OPTIONS(path string, handler Handler) Route
	HEAD(path string, handler Handler) Route
	Route(methods []string, path string, handler Handler) Route
}
