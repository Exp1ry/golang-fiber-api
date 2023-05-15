package routes

import (
	"api.ainvest.com/controller/api/handlers"
	stockBroker "api.ainvest.com/controller/pkg/stocks"
	"github.com/gofiber/fiber/v2"
)

func StockRoutes(api fiber.Router, service stockBroker.Service){
api.Get("/stock/all", handlers.GetAllStockBrokers(service))
api.Post("/stock/new", handlers.AddNewStockBroker(service))
api.Put("/stock/edit", handlers.UpdateOneStockBroker(service))
api.Delete("/stock/remove", handlers.DeleteStockBroker(service))
}