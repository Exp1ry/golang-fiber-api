package handlers

import (
	"errors"

	"api.ainvest.com/controller/api/presenters"
	"api.ainvest.com/controller/models"
	nftBroker "api.ainvest.com/controller/pkg/nft"
	"github.com/gofiber/fiber/v2"
)

func GetAllNftBrokers(service nftBroker.Service) fiber.Handler {
	return func(c *fiber.Ctx) error{
		resp, err := service.GetAllNFTBrokers()
		if err!= nil{
			return c.Status(401).JSON(presenters.DynamicResponse(map[string]string{}, "",  err, true))
		}

		return c.Status(200).JSON(presenters.DynamicResponse(resp, "Successfully fetched Nft brokers.",  nil, false))
	}
}

func UpdateOneNftBroker(service nftBroker.Service) fiber.Handler {

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

		ok, err := service.UpdateNFTBroker(body.ID, body.Update)
if err!= nil {
	return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{}, "",  err, true))

}
		return c.Status(200).JSON(presenters.DynamicResponse(map[string]bool{"success":ok},"Successfully updated Nft broker.",  nil, false))

	}
}

func AddNewNftBroker(service nftBroker.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body struct {
			ID string `json:"_id"`
			Broker *models.NFTBrokerModel `json:"broker"`
		}

		err := c.BodyParser(&body)
		if err!= nil {
			return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{},"",   err, true))

		}

		if body.Broker == nil {
			return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{}, "",  errors.New("Please pass the correct parameters"), true))

		}

		err = service.InsertNewBroker(body.Broker)
		if err!= nil {
			return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{}, "",  err, true))

		}

		return c.Status(200).JSON(presenters.DynamicResponse(map[string]string{"success":"true"},"Successfully updated Nft broker.",  nil, false))

	}
}

func DeleteNftBroker(service nftBroker.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body struct {
			ID string `json:_id"`
		}

		err := c.BodyParser(&body)
		if err!= nil {
			return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{},  "", err, true))

		}
		
		if body.ID == ""{
			return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{}, "",  errors.New("Please pass the correct parameters"), true))

		}

		ok, err := service.DeleteNFTBroker(body.ID)
		if err!= nil {
			return c.Status(400).JSON(presenters.DynamicResponse(map[string]string{},  "", err, true))

		}
		return c.Status(200).JSON(presenters.DynamicResponse(map[string]bool{"success":ok},"Successfully updated Nft broker.",  nil, false))

	}
}