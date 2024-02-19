package arc

import "net/http"

const (
	MethodGet     = "GET"
	MethodHead    = "HEAD"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodPatch   = "PATCH" // RFC 5789
	MethodDelete  = "DELETE"
	MethodConnect = "CONNECT"
	MethodOptions = "OPTIONS"
	MethodTrace   = "TRACE"
)

type Usable interface {
	HandlerFunc | Router
}

type Map map[string]any

type ParamFunc func(*Router)
type HandlerFunc func(*Ctx) error

type Ctx struct {
	Response http.ResponseWriter
	Request  *http.Request
	local    map[string]any
}

type RouterConfig struct {
	Prefix  string
	Logging bool
}

type Router struct {
	Config     RouterConfig
	handlers   []*Handler
	middleware []*Middleware
}

// basic route handler
type Handler struct {
	Route  string
	Method string
	Func   HandlerFunc
}

// middleware
type Middleware struct {
	Prefix string
	Func   HandlerFunc
}
