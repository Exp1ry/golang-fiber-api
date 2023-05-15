package routes

import (
	"api.ainvest.com/controller/api/handlers"
	nftBroker "api.ainvest.com/controller/pkg/nft"
	"github.com/gofiber/fiber/v2"
)


func NftRoutes(api fiber.Router, service nftBroker.Service){
	api.Get("/nft/all", handlers.GetAllNftBrokers(service))
	api.Post("/nft/new", handlers.AddNewNftBroker(service))
	api.Put("/nft/edit", handlers.UpdateOneNftBroker(service))
	api.Delete("/nft/remove", handlers.DeleteNftBroker(service))
}