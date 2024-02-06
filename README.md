# Arc
Arc, initially, is just a router built on top of net/http, but it is designed to be batteries included and provide complete functionality of a web framework.
This is my learning project which I hope to complete. I previously had a simple framework named Cyberia, but later on its structure and implementation seemed too strange to me.
Just like Cyberia, this is heavily inspired by Fiber.

# Usage
Arc is built purely on top of net/http, so it has zero dependencies.
```
go get github.com/intervinn/arc
```

```go
package main

import "github.com/intervinn/arc"

func main() {
	router := arc.NewRouter(arc.WithLogging())
	router.GET("/", func(c *arc.Ctx) error {
		return c.JSON(arc.StatusOK, arc.Map{
			"message": "hello arc!",
		})
	})

	router.Listen(":8080")
}
```

Check examples folder for more.

# TODO
* Dynamic Routing
* Route Params
* Templates?