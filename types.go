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

type Map map[string]any

type ParamFunc func(*Router)
type HandlerFunc func(*Ctx) error

type Ctx struct {
	Response http.ResponseWriter
	Request  *http.Request
}

/*
DYNAMIC ROUTING
the route will have routes starting with :
parse the route template and copy the param name and its index
parse the url and return param values at index

PARSING URL
first see static urls

then see dynamic urls, check if slash splits are equal and
values parse properly

if any error happens just continue and if it doenst find anything do something

*/

type RouterConfig struct {
	Prefix  string
	Logging bool
}

type Router struct {
	Config   RouterConfig
	Handlers []*Handler
}

type Handler struct {
	Route  string
	Method string
	Func   HandlerFunc
}
