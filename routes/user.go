package routes

import (
	"fiber/backend/handler"
	"fiber/backend/middleware"

	"github.com/gofiber/fiber/v2"
)




func RouteUser(route *fiber.App) {

	api :=route.Group("/api/v1")
	api.Use(middleware.MiddlewareAccess)
	api.Get("/user",handler.GetAllUser)
	api.Get("/user/:id",middleware.AdminAcces,handler.GetById)
	api.Post("/user",middleware.AdminAcces,handler.UserCreate)
	api.Patch("/user/:id",middleware.AdminAcces,handler.UserUpdate)
	api.Delete("/user/:id",middleware.AdminAcces,handler.DeleteUser)
}