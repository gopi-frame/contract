package router

import (
	"net/http"

	rc "github.com/gopi-frame/contract/response"
)

type Middleware interface {
	Handle(*http.Request, func(*http.Request) rc.Responser) rc.Responser
}

type ConstructableMiddleware interface {
	Middleware
	Construct(request *http.Request)
}
