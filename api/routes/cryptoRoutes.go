package routes

import (
	"api.ainvest.com/controller/api/handlers"
	cryptoBrokers "api.ainvest.com/controller/pkg/crypto"
	"github.com/gofiber/fiber/v2"
)

func CryptoRoutes(api fiber.Router, service cryptoBrokers.Service){
	api.Get("/crypto/all", handlers.GetAllCryptoBrokers(service))
	api.Post("/crypto/new", handlers.AddNewCryptoBroker(service))
	api.Put("/crypto/edit", handlers.UpdateOneCryptoBroker(service))
	api.Delete("/crypto/remove", handlers.DeleteCryptoBroker(service))
}