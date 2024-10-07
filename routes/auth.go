package routes

import (
	"fiber/backend/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteAuth(route *fiber.App) {
	route.Post("/auth-create", handler.AuthCreate)
	route.Post("/login", handler.Login)
	route.Post("/refresh-token", handler.RefreshToken)
}