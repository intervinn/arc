package main

import (
	"fmt"

	"github.com/intervinn/arc"
)

func main() {
	tree := arc.NewTree()

	r := "/v1/users/intervinn"
	p := arc.Pieces(r)
	tree.Add(p, arc.NewHandler(r, func(c *arc.Context) error { return nil }))

	fmt.Println(tree.Get(p).Route)
}
