package middleware

import (
	"fmt"
	"strings"

	"api.ainvest.com/controller/api/presenters"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func ValidateToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		cookie := c.Cookies("token")

		if cookie == ""{
			return c.SendStatus(200)
		}
		fmt.Println(cookie)

		tknStr := strings.Replace(cookie, "Bearer ", "", 1)
		token,err := jwt.Parse(tknStr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
			}
			return []byte("secret"), nil
		})
		
		if err!= nil {
			fmt.Println(err)
			if err, ok := err.(*fiber.Error); ok {
				return err
			}
			return c.Status(fiber.StatusUnauthorized).JSON(presenters.DynamicResponse(fiber.Map{}, "Invalid token", err, true))
		}

		if !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(presenters.DynamicResponse(map[string]string{}, "Invalid token", nil, true))
		}

		c.Locals("token", token)

		return c.Next()
	}
}