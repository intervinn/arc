package arc

type HandlerFunc func(c *Context) error

type Handler struct {
	Route  string
	Method string
	Call   HandlerFunc
}

func NewHandler(method string, route string, call HandlerFunc) *Handler {
	return &Handler{
		Route: route,
		Call:  call,
	}
}
