package routes

import (
	"api.ainvest.com/controller/api/handlers"
	forexBrokers "api.ainvest.com/controller/pkg/forex"
	"github.com/gofiber/fiber/v2"
)

func ForexRoutes(api fiber.Router, service forexBrokers.Service){
	api.Get("/all", handlers.GetAllForexBrokers(service))
	api.Post("/new", handlers.AddNewForexBroker(service))
	api.Put("/edit", handlers.UpdateOneForexBroker(service))
	api.Delete("/remove", handlers.DeleteForexBroker(service))
}