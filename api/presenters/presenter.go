package presenters

import "github.com/gofiber/fiber/v2"


func DynamicResponse(data interface{}, msg string,  err error, errBool bool) *fiber.Map {
	return &fiber.Map{
		"data":          data,
		"message":       msg,
		"errorMessage": err,
		"error":         errBool,
	}
}