package arc

import "net/http"

type Context struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Request:        r,
		ResponseWriter: w,
	}
}
