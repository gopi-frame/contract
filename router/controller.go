package router

import "net/http"

type Controller interface {
	RouterGroup() RouterGroup
}

type ConstructableController interface {
	Controller
	Construct(*http.Request)
}
