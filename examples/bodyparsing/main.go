package main

import "github.com/intervinn/arc"

type HelloBody struct {
	Message string `json:"message"`
}

func main() {
	router := arc.NewRouter(arc.WithLogging())
	router.POST("/hello", func(c *arc.Ctx) error {

		body := new(HelloBody)
		if err := c.BodyParse(body); err != nil {
			return err
		}

		return c.JSON(arc.StatusOK, arc.Map{
			"you said": body.Message,
		})
	})

	router.Listen(":8080")
}
