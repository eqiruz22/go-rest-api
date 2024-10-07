package middleware

import (
	"fiber/backend/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func MiddlewareAccess(c *fiber.Ctx) error {
    authHeader := c.Get("Authorization")

	// if authorization not include for request it will throw unauthorized
	// or if remove this code it will throw panic!!
    if len(authHeader) == 0 {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "message": "Authorization header is required",
        })
    }
    if !strings.HasPrefix(authHeader, "Bearer ") {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "message": "Authorization header must start with Bearer",
        })
    }
    tokenString := authHeader[len("Bearer "):]
    claims, err := utils.DecodeToken(tokenString)
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "message": "invalid token",
        })
    }
    c.Locals("userInfo", claims)
    return c.Next()
}
