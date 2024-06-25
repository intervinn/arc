package arc

type TreeNode struct {
	Handler *Handler

	Key string // For splitted routes
	Sub []*TreeNode
}

func NewTree() *TreeNode {
	return &TreeNode{
		Handler: nil,
		Key:     "",
		Sub:     []*TreeNode{},
	}
}

func (t *TreeNode) search(key string) *TreeNode {
	for _, v := range t.Sub {
		if v.Key == key {
			return v
		}
	}
	return nil
}

// ex: /hello
func (t *TreeNode) Add(p []string, h *Handler) {
	node := t
	len := len(p)

	for i := 0; i < len-1; i++ {
		v := p[i]
		n := node.search(v)
		if n == nil {
			n = &TreeNode{
				Handler: nil,
				Key:     v,
				Sub:     []*TreeNode{},
			}
			node.Sub = append(node.Sub, n)
		}
		node = n
	}

	node.Sub = append(node.Sub, &TreeNode{
		Handler: h,
		Key:     p[len-1],
		Sub:     []*TreeNode{},
	})
}

func (t *TreeNode) Get(p []string) *Handler {
	node := t

	for _, v := range p {
		n := node.search(v)
		if n == nil {
			return nil
		}

		node = n
	}
	return node.Handler
}
