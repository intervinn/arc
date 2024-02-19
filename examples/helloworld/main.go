package main

import "github.com/intervinn/arc"

func main() {

	cats := arc.NewRouter(arc.WithPrefix("/cats"))
	cats.GET("/marta", func(c *arc.Ctx) error {
		return c.JSON(arc.StatusOK, arc.Map{
			"marta": "is a very nice cat",
		})
	})

	router := arc.NewRouter(arc.WithLogging())
	router.GET("/", func(c *arc.Ctx) error {
		return c.JSON(arc.StatusOK, arc.Map{
			"message": "hello arc!",
		})
	})

	router.Mount(cats)

	router.Listen(":8080")
}
