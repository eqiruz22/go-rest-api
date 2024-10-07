package routes

import (
	"fiber/backend/handler"
	"fiber/backend/middleware"

	"github.com/gofiber/fiber/v2"
)




func RouteUser(route *fiber.App) {

	route.Get("/user",handler.GetAllUser)
	//route.Use()
	route.Get("/user/:id",middleware.MiddlewareAccess,handler.GetById)
	route.Post("/user",middleware.MiddlewareAccess,handler.UserCreate)
	route.Patch("/user/:id",middleware.MiddlewareAccess,handler.UserUpdate)
	route.Delete("/user/:id",middleware.MiddlewareAccess,handler.DeleteUser)
}