package router

import (
	"net/http"

	"github.com/gopi-frame/contract/response"
)

type Router interface {
	http.Handler
	Use(...Middleware)
	Group(RouteGroup, func(Router)) Router
	GET(path string, handler func(request *http.Request) response.Responser) Route
	POST(path string, handler func(request *http.Request) response.Responser) Route
	PUT(path string, handler func(request *http.Request) response.Responser) Route
	PATCH(path string, handler func(request *http.Request) response.Responser) Route
	DELETE(path string, handler func(request *http.Request) response.Responser) Route
	OPTIONS(path string, handler func(request *http.Request) response.Responser) Route
	Route(methods []string, path string, handler func(request *http.Request) response.Responser) Route
}
