package arc

import (
	"log"
	"net/http"
)

// router params
func WithPrefix(prefix string) ParamFunc {
	return func(r *Router) {
		r.Config.Prefix = prefix
	}
}

func WithLogging() ParamFunc {
	return func(r *Router) {
		r.Config.Logging = true
	}
}

// router implementation
func NewRouter(params ...ParamFunc) *Router {

	defcnf := RouterConfig{
		Prefix:  "",
		Logging: false,
	}

	r := &Router{
		Config:   defcnf,
		Handlers: []*Handler{},
	}

	for _, v := range params {
		v(r)
	}

	return r
}

// route registering
func (r *Router) Route(route string, method string, hfunc HandlerFunc) {
	h := &Handler{
		Route:  ConnectRoutes(r.Config.Prefix, route),
		Method: method,
		Func:   hfunc,
	}

	if r.Config.Logging {
		log.Printf("bound route %s\n", h.Route)
	}

	r.Handlers = append(r.Handlers, h)
}

// route searching
func (r *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	ctx := &Ctx{
		Request:  req,
		Response: res,
	}

	for _, v := range r.Handlers {
		if v.Route == req.URL.Path {
			err := v.Func(ctx)
			if err != nil {
				panic(err)
			}
		}
	}
}

// registering for different methods
func (r *Router) POST(route string, hfunc HandlerFunc) {
	r.Route(route, MethodPost, hfunc)
}
func (r *Router) GET(route string, hfunc HandlerFunc) {
	r.Route(route, MethodGet, hfunc)
}
func (r *Router) OPTIONS(route string, hfunc HandlerFunc) {
	r.Route(route, MethodOptions, hfunc)
}
func (r *Router) DELETE(route string, hfunc HandlerFunc) {
	r.Route(route, MethodDelete, hfunc)
}
func (r *Router) HEAD(route string, hfunc HandlerFunc) {
	r.Route(route, MethodHead, hfunc)
}
func (r *Router) PATCH(route string, hfunc HandlerFunc) {
	r.Route(route, MethodPatch, hfunc)
}
func (r *Router) PUT(route string, hfunc HandlerFunc) {
	r.Route(route, MethodPut, hfunc)
}
func (r *Router) TRACE(route string, hfunc HandlerFunc) {
	r.Route(route, MethodTrace, hfunc)
}
func (r *Router) CONNECT(route string, hfunc HandlerFunc) {
	r.Route(route, MethodConnect, hfunc)
}

func (r *Router) Listen(addr string) {

	if r.Config.Logging {
		log.Printf("server listening on address %s", addr)
	}

	http.ListenAndServe(addr, r)
}
