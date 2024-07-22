package router

import "net/http"

type Controller interface {
	RouteGroup() RouteGroup
}

type ConstructableController interface {
	Controller
	Construct(*http.Request)
}
