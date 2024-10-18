package routes

import (
	"fiber/backend/handler"
	"fiber/backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func RouteAuth(route *fiber.App) {
	api := route.Group("/api/v1")
	api.Post("/login", handler.Login)
	api.Use(middleware.MiddlewareAccess)
	api.Post("/auth-create", middleware.AdminAcces,handler.AuthCreate)
	api.Post("/refresh-token", handler.RefreshToken)
}