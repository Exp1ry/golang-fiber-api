package handlers

import (
	"errors"
	"fmt"

	"api.ainvest.com/controller/api/presenters"
	"api.ainvest.com/controller/models"
	forexBrokers "api.ainvest.com/controller/pkg/forex"
	"github.com/gofiber/fiber/v2"
)

func GetAllForexBrokers(service forexBrokers.Service) fiber.Handler {
	return func(c *fiber.Ctx) error{
		resp, err := service.GetAllForexBrokers()
		if err!= nil{
			return c.Status(401).JSON(presenters.DynamicResponse(map[string]string{}, "", err, true))
		}

		return c.Status(200).JSON(presenters.DynamicResponse(resp, "Successfully fetched Forex brokers.",  nil, false))
	}
}

func UpdateOneForexBroker(service forexBrokers.Service) fiber.Handler {

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

		ok, err := service.UpdateForexBroker(body.ID, body.Update)
if err!= nil {
	return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{}, "", err, true))

}
		return c.Status(200).JSON(presenters.DynamicResponse(map[string]bool{"success":ok},"Successfully updated Forex broker.", nil, false))

	}
}

func AddNewForexBroker(service forexBrokers.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body struct {
			ID string `json:"_id"`
			Broker *models.ForexBrokerModel `json:"broker"`
		}

		err := c.BodyParser(&body)
		if err!= nil {
			return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{}, "", err, true))

		}

		if body.Broker == nil {
			return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{}, "", errors.New("Please pass the correct parameters"), true))

		}

		err = service.InsertNewForexBroker(body.Broker)
		if err!= nil {
			return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{}, "", err, true))

		}

		return c.Status(200).JSON(presenters.DynamicResponse(map[string]string{"success":"true"},"Successfully updated Forex broker.", nil, false))

	}
}

func DeleteForexBroker(service forexBrokers.Service) fiber.Handler {
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

		ok, funcErr := service.DeleteForexBroker(body.ID)
		if err!= nil {
			return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{}, "", funcErr, true))

		}
		return c.Status(200).JSON(presenters.DynamicResponse(map[string]bool{"success":ok},"Successfully updated Forex broker.",  nil, false))

	}
}