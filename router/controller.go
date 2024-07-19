package router

import "net/http"

type Controller interface {
	RouterGroup() RouteGroup
}

type ConstructableController interface {
	Controller
	Construct(*http.Request)
}
