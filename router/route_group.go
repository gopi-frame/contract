package router

type RouteGroup interface {
	Build() Router
}
