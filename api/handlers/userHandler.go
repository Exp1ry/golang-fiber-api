package handlers

import (
	"errors"
	"fmt"
	"time"

	"api.ainvest.com/controller/api/presenters"
	"api.ainvest.com/controller/models"
	"api.ainvest.com/controller/pkg/users"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetAllUsersHandler(service users.Service) fiber.Handler{
	return func(c *fiber.Ctx) error {

		token, ok := c.Locals("token").(*jwt.Token)
		if ok {
			
			claims, ok := token.Claims.(jwt.MapClaims)
			if ok {

				email := claims["email"].(string)
				fmt.Println(email)
			}
		}
		resp, err := service.FetchAllUsers()
		if err!= nil {
			return c.JSON(presenters.DynamicResponse(map[string]string{}, "", err, true))
		}

		


		return c.JSON(presenters.DynamicResponse(resp, "Successfully fetched users", nil, false))
	}
}

func UpdateOneUser(service users.Service) fiber.Handler {

	return func(c *fiber.Ctx) error {
var body struct {
	ID string `json:"_id"`
	Update map[string]interface{} `json:"update"`
}
err := c.BodyParser(&body)
if err!= nil {
	return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{}, "", err, true))
}

if body.ID == "" {
	return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{}, "", errors.New("Please pass the correct parameters"), true))
	
}
fmt.Println(body.Update)

		ok, err := service.UpdateUser(body.ID, body.Update)
if err!= nil {
	return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{}, "", err, true))

}
		return c.Status(200).JSON(presenters.DynamicResponse(map[string]bool{"success":ok},"Successfully updated Forex broker.", nil, false))

	}
}

func AddNewUser(service users.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body struct {
			ID string `json:"_id"`
			User *models.UserModel `json:"broker"`
		}

		err := c.BodyParser(&body)
		if err!= nil {
			fmt.Println(err)
			return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{}, "", err, true))

		}

		if body.User == nil {
			fmt.Println(err)

			return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{}, "", errors.New("Please pass the correct parameters"), true))

		}

		err = service.AddUser(body.User)
		if err!= nil {
			fmt.Println(err)

			return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{}, "", err, true))

		}

		return c.Status(200).JSON(presenters.DynamicResponse(map[string]string{"success":"true"},"Successfully updated Forex broker.", nil, false))

	}
}

func DeleteOneUser(service users.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body struct {
			ID string `json:"_id"`
		}

		err := c.BodyParser(&body)
		if err!= nil {
			return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{}, "", err, true))

		}
		
		if body.ID == ""{
			return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{}, "", errors.New("Please pass the correct parameters"), true))

		}

		ok, funcErr := service.DeleteUser(body.ID)
		if err!= nil {
			return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{}, "", funcErr, true))

		}
		return c.Status(200).JSON(presenters.DynamicResponse(map[string]bool{"success":ok},"Successfully updated Forex broker.",  nil, false))

	}
}

func HandleAdminSignup(service users.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body struct {
			Email string `json:"email"`
			Password string `json:"password"`
			FirstName string `json:"first_name"`
			LastName string `json:"last_name"`
		}

		err := c.BodyParser(&body)
		if err!= nil {
			return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{}, "Please pass the correct params",err, true ))
		}
fmt.Println(body)
		ok, err := service.AddAdmin(body.Email, body.Password, body.FirstName,body.LastName)
		if err!= nil {
			fmt.Println(err)

			return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{}, "Error signing up", err, ok))
		}

		return c.Status(200).JSON(presenters.DynamicResponse(fiber.Map{"success":ok}, "Successfully signed up",nil, ok))
	}
}

func HandleAdminSignin(service users.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
	var body struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}

		err := c.BodyParser(&body)
		if err!= nil {
			return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{}, "Please pass the correct params",err, true ))
		}

		signedToken, err := service.EnterAdmin(body.Email, body.Password)
		if err!= nil {
			return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{}, "Error signing up", err, true))
		}



// Create a new cookie
cookie := fiber.Cookie{
	Name:     "token",
	Value:    signedToken,
	Expires:  time.Now().Add(time.Hour * 24), // Set the cookie expiration time (e.g., 24 hours)
	HTTPOnly: false,                           // Set HTTPOnly to true for better security
}

// Set the cookie in the response
c.Cookie(&cookie)
		return c.Status(200).JSON(presenters.DynamicResponse(cookie, "Successfully signed up",nil, false))
	}
}


func CheckAuth(c *fiber.Ctx) error {
	// Retrieve the token from the cookie
	cookie := c.Cookies("token")
	if cookie == "" {
		// Cookie is missing, user is not authenticated
		return c.JSON(fiber.Map{
			"authenticated": false,
		})
	}

	// Verify and parse the token
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		// Verify the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Replace "secret" with your actual secret key used to sign the token
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid {
		// Token is invalid, user is not authenticated
		return c.JSON(fiber.Map{
			"authenticated": false,
		})
	}

	// Token is valid, user is authenticated
	return c.JSON(fiber.Map{
		"authenticated": true,
	})
}
