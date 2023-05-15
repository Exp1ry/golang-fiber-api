package handlers

import (
	"errors"
"fmt"
	"api.ainvest.com/controller/api/presenters"
	"api.ainvest.com/controller/models"
	cryptoBrokers "api.ainvest.com/controller/pkg/crypto"
	"github.com/gofiber/fiber/v2"
)



func GetAllCryptoBrokers(service cryptoBrokers.Service) fiber.Handler {
	return func(c *fiber.Ctx) error{
		resp, err := service.GetAllCryptoBrokers()
		if err!= nil{
			fmt.Println(err)
			return c.Status(401).JSON(presenters.DynamicResponse(map[string]string{},  "", err, true))
		}

		return c.Status(200).JSON(presenters.DynamicResponse(resp, "Successfully fetched crypto brokers.",  nil, false))
	}
}

func UpdateOneCryptoBroker(service cryptoBrokers.Service) fiber.Handler {

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
	return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{},  "", errors.New("Please pass the correct parameters"), true))

}

		ok, err := service.UpdateCryptoBroker(body.ID, body.Update)
if err!= nil {
	return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{}, "",  err, true))

}
		return c.Status(200).JSON(presenters.DynamicResponse(map[string]bool{"success":ok},"Successfully updated crypto broker.",  nil, false))

	}
}

func AddNewCryptoBroker(service cryptoBrokers.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body struct {
			ID string `json:"_id"`
			Broker *models.CryptoBrokerModel `json:"broker"`
		}

		err := c.BodyParser(&body)
		if err!= nil {
			return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{}, "",  err, true))

		}

		if body.Broker == nil {
			return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{}, "",  errors.New("Please pass the correct parameters"), true))

		}

		err = service.InsertNewBroker(body.Broker)
		if err!= nil {
			return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{}, "", err, true))

		}

		return c.Status(200).JSON(presenters.DynamicResponse(map[string]string{"success":"true"},"Successfully updated crypto broker.",  nil, false))

	}
}

func DeleteCryptoBroker(service cryptoBrokers.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body struct {
			ID string `json:"_id"`
		}

		err := c.BodyParser(&body)
		if err!= nil {
			return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{}, err.Error(),  err, true))

		}
		
		if body.ID == ""{
			return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{},"",   errors.New("Please pass the correct parameters"), true))

		}

		ok, err := service.DeleteBroker(body.ID)
		if err!= nil {
			return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{},  "", err, true))

		}
		return c.Status(200).JSON(presenters.DynamicResponse(map[string]bool{"success":ok},"Successfully updated crypto broker.",  nil, false))

	}
}