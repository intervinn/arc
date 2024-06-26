package arc

import "net/http"

type Router struct {
	Tree *Node
}

func NewRouter() *Router {
	return &Router{
		Tree: NewNode("", nil, []*Node{}),
	}
}

func (rt *Router) Route(h *Handler) {
	rt.Tree.Add(Pieces(h.Route), h)
}

func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pieces := Pieces(r.URL.Path)
	h := rt.Tree.Get(pieces, r.Method)
	ctx := NewContext(w, r)

	if h != nil {
		h.Call(ctx)
	}
}
