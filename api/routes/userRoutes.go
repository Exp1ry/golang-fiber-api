package routes

import (
	"api.ainvest.com/controller/api/handlers"
	"api.ainvest.com/controller/pkg/users"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(api fiber.Router, service users.Service){
api.Get("/all", handlers.GetAllUsersHandler(service))
api.Post("/new", handlers.AddNewUser(service))
api.Put("/edit", handlers.UpdateOneUser(service))
api.Delete("/remove", handlers.DeleteOneUser(service))

api.Post("/admin/signup", handlers.HandleAdminSignup(service))
api.Post("/admin/signin", handlers.HandleAdminSignin(service))

api.Get("/auth", handlers.CheckAuth)
}

func AdminRoutes(api fiber.Router, service users.Service) {
	api.Post("/signup", handlers.HandleAdminSignup(service))
api.Post("/signin", handlers.HandleAdminSignin(service))
}

