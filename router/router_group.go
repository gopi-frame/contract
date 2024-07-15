package router

type RouterGroup interface {
	Build() Router
}
