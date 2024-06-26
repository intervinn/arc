package main

import (
	"net/http"

	"github.com/intervinn/arc"
)

func main() {
	r := arc.NewRouter()

	http.ListenAndServe(":8080", r)
}
