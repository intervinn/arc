package arc

type HandlerFunc func(c *Context) error

type Handler struct {
	Route string
	Call  HandlerFunc
}

func NewHandler(route string, call HandlerFunc) *Handler {
	return &Handler{
		Route: route,
		Call:  call,
	}
}
