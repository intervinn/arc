package arc

type Node struct {
	Handler *Handler
	Pattern string
	Key     string // For splitted routes
	Sub     []*Node
}

func NewNode(field string, handler *Handler, sub []*Node) *Node {
	res := &Node{
		Handler: handler,
		Sub:     []*Node{},
	}

	if IsPattern(field) {
		res.Pattern = field[1:]
	} else {
		res.Key = field
	}

	return res
}

func (t *Node) search(key string, method string) *Node {
	for _, v := range t.Sub {
		if v.Key == key && v.Handler.Method == method {
			return v
		}
	}

	// if theres no exact keys, let's search including patterns
	for _, v := range t.Sub {
		if !IsEmpty(v.Pattern) && v.Handler.Method == method {
			return v
		}
	}
	return nil
}

// Adds a new handler to the tree, creating new nodes if these don't exist
func (t *Node) Add(p []string, h *Handler) {
	node := t
	len := len(p)

	for i := 0; i < len-1; i++ {
		v := p[i]
		n := node.search(v, h.Method)
		if n == nil {
			n = NewNode(v, nil, []*Node{})
			node.Sub = append(node.Sub, n)
		}
		node = n
	}

	node.Sub = append(node.Sub, NewNode(p[len-1], h, []*Node{}))
}

func (t *Node) Get(p []string, method string) *Handler {
	node := t

	for _, v := range p {
		n := node.search(v, method)
		if n == nil {
			return nil
		}
		node = n
	}
	return node.Handler
}
