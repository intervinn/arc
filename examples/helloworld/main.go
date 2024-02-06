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
