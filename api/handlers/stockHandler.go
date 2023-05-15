package handlers

import (
	"errors"

	"api.ainvest.com/controller/api/presenters"
	"api.ainvest.com/controller/models"
	stockBroker "api.ainvest.com/controller/pkg/stocks"

	"github.com/gofiber/fiber/v2"
)

func GetAllStockBrokers(service stockBroker.Service) fiber.Handler {
	return func(c *fiber.Ctx) error{
		resp, err := service.GetAllStockBrokers()
		if err!= nil{
			return c.Status(401).JSON(presenters.DynamicResponse(map[string]string{},  "Error executing function.", err, true))
		}

		return c.Status(200).JSON(presenters.DynamicResponse(resp, "Successfully fetched Stock brokers.",  nil, false))
	}
}

func UpdateOneStockBroker(service stockBroker.Service) fiber.Handler {

	return func(c *fiber.Ctx) error {
var body struct {
	ID string `json:"_id"`
	Update map[string]interface{} `json:"update"`
}

err := c.BodyParser(&body)
if err!= nil {
	return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{}, "",  err, true))
}

if body.ID == "" {
	return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{}, "",  errors.New("Please pass the correct parameters"), true))

}

		ok, err := service.UpdateStockBroker(body.ID, body.Update)
if err!= nil {
	return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{},  "", err, true))

}
		return c.Status(200).JSON(presenters.DynamicResponse(map[string]bool{"success":ok},"Successfully updated Stock broker.",  nil, false))

	}
}

func AddNewStockBroker(service stockBroker.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body struct {
			ID string `json:"_id"`
			Broker *models.StockBrokerModel `json:"broker"`
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
			return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{}, "",  err, true))

		}

		return c.Status(200).JSON(presenters.DynamicResponse(map[string]string{"success":"true"},"Successfully updated Stock broker.",  nil, false))

	}
}

func DeleteStockBroker(service stockBroker.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body struct {
			ID string `json:_id"`
		}

		err := c.BodyParser(&body)
		if err!= nil {
			return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{}, "",  err, true))

		}
		
		if body.ID == ""{
			return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{}, "",  errors.New("Please pass the correct parameters"), true))

		}

		ok, err := service.DeleteStockBroker(body.ID)
		if err!= nil {
			return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{}, "",  err, true))

		}
		return c.Status(200).JSON(presenters.DynamicResponse(map[string]bool{"success":ok},"Successfully updated Stock broker.",  nil, false))

	}
}