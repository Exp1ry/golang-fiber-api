package routes

import (
	"api.ainvest.com/controller/api/handlers"
	cryptoBrokers "api.ainvest.com/controller/pkg/crypto"
	"github.com/gofiber/fiber/v2"
)

func CryptoRoutes(api fiber.Router, service cryptoBrokers.Service){
	api.Get("/all", handlers.GetAllCryptoBrokers(service))
	api.Post("/new", handlers.AddNewCryptoBroker(service))
	api.Put("/edit", handlers.UpdateOneCryptoBroker(service))
	api.Delete("/remove", handlers.DeleteCryptoBroker(service))
}