package routes

import (
	"api.ainvest.com/controller/api/handlers"
	nftBroker "api.ainvest.com/controller/pkg/nft"
	"github.com/gofiber/fiber/v2"
)


func NftRoutes(api fiber.Router, service nftBroker.Service){
	api.Get("/all", handlers.GetAllNftBrokers(service))
	api.Post("/new", handlers.AddNewNftBroker(service))
	api.Put("/edit", handlers.UpdateOneNftBroker(service))
	api.Delete("/remove", handlers.DeleteNftBroker(service))
}