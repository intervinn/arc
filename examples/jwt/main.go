package main

import (
	"errors"

	"github.com/golang-jwt/jwt"
)

type User struct {
	Name string `json:"name"`
}

func CreateToken(u User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": u.Name,
	})

	str, err := token.SignedString([]byte("epic"))
	if err != nil {
		panic(err)
	}
	return str
}

func VerifyToken(str string) (string, error) {
	token, err := jwt.Parse(str, func(t *jwt.Token) (interface{}, error) {
		return []byte("epic"), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims["name"].(string), nil
	}
	return "", errors.New("invalid token")
}

/*
func CheckAuth(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	user, err := VerifyToken(token)
	if err != nil {
		return c.Next()
	}

	fmt.Printf("token %s valid for %s", token, user)

	return c.Next()
}

func main() {
	users := []User{}
	app := fiber.New()

	app.Use(CheckAuth)

	app.Post("/auth/register", func(c *fiber.Ctx) error {
		u := new(User)
		if err := c.BodyParser(u); err != nil {
			return err
		}

		users = append(users, *u)
		token := CreateToken(*u)

		return c.JSON(fiber.Map{
			"token": token,
		})
	})

	app.Listen(":8080")
}
*/
