package websocket

// Handler is the interface for the websocket handler.
type Handler interface {
	// Handle handles the websocket message.
	Handle(client Client, message Message)
}

type ErrorHandler interface {
	Handle(client Client, err any)
}

// Router is the interface for the websocket router.
type Router interface {
	// Use adds middlewares to the router.
	Use(middlewares ...Middleware)
	// On registers a handler for a specific event.
	On(event string, handler Handler) Route
	// Serve serves the websocket connection.
	Serve(client Client, message Message)
	// OnNotFound registers a handler for not found events.
	OnNotFound(handler Handler)
	// OnError registers a handler for errors.
	OnError(errorHandler ErrorHandler)
}

// Route is the interface for the websocket route.
type Route interface {
	// Use adds middlewares to the route.
	Use(middlewares ...Middleware)
}
