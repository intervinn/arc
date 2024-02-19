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

// Create a new router
func NewRouter(params ...ParamFunc) *Router {

	defcnf := RouterConfig{
		Prefix:  "",
		Logging: false,
	}

	r := &Router{
		Config:   defcnf,
		handlers: []*Handler{},
	}

	for _, v := range params {
		v(r)
	}

	return r
}

// Registers a route for specific route and method
func (r *Router) Route(route string, method string, hfunc HandlerFunc) {
	h := &Handler{
		Route:  ConnectRoutes(r.Config.Prefix, route),
		Method: method,
		Func:   hfunc,
	}

	if r.Config.Logging {
		log.Printf("bound route %s\n", h.Route)
	}

	r.handlers = append(r.handlers, h)
}

func (r *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	ctx := &Ctx{
		Request:  req,
		Response: res,
	}

	for _, v := range r.middleware {
		err := v.Func(ctx)
		if err != nil {
			panic(err)
		}
	}

	for _, v := range r.handlers {
		if v.Route == req.URL.Path {
			err := v.Func(ctx)
			if err != nil {
				panic(err)
			}
		}
	}
}

// Calls Route() with POST Method
func (r *Router) POST(route string, hfunc HandlerFunc) {
	r.Route(route, MethodPost, hfunc)
}

// Calls Route() with GET Method
func (r *Router) GET(route string, hfunc HandlerFunc) {
	r.Route(route, MethodGet, hfunc)
}

// Calls Route() with OPTIONS Method
func (r *Router) OPTIONS(route string, hfunc HandlerFunc) {
	r.Route(route, MethodOptions, hfunc)
}

// Calls Route() with DELETE Method
func (r *Router) DELETE(route string, hfunc HandlerFunc) {
	r.Route(route, MethodDelete, hfunc)
}

// Calls Route() with HEAD Method
func (r *Router) HEAD(route string, hfunc HandlerFunc) {
	r.Route(route, MethodHead, hfunc)
}

// Calls Route() with PATCH Method
func (r *Router) PATCH(route string, hfunc HandlerFunc) {
	r.Route(route, MethodPatch, hfunc)
}

// Calls Route() with PUT Method
func (r *Router) PUT(route string, hfunc HandlerFunc) {
	r.Route(route, MethodPut, hfunc)
}

// Calls Route() with TRACE Method
func (r *Router) TRACE(route string, hfunc HandlerFunc) {
	r.Route(route, MethodTrace, hfunc)
}

// Calls Route() with CONNECT Method
func (r *Router) CONNECT(route string, hfunc HandlerFunc) {
	r.Route(route, MethodConnect, hfunc)
}

// Merges other router's handlers and middleware
func (r *Router) Mount(other *Router) {
	for _, v := range other.handlers {
		r.Route(v.Route, v.Method, v.Func)
	}

	for _, v := range other.middleware {
		nware := &Middleware{
			Prefix: ConnectRoutes(r.Config.Prefix, other.Config.Prefix, v.Prefix),
			Func:   v.Func,
		}

		r.middleware = append(r.middleware, nware)
	}
}

// Add a middleware
// Put empty string in first argument to handle all routes
func (r *Router) Use(prefix string, mfunc HandlerFunc) {
	r.middleware = append(r.middleware, &Middleware{
		Prefix: prefix,
		Func:   mfunc,
	})
}

// Start the server and listen at address `addr`
func (r *Router) Listen(addr string) {

	if r.Config.Logging {
		log.Printf("server listening on address %s", addr)
	}

	http.ListenAndServe(addr, r)
}
