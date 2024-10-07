package utils

import "github.com/gofiber/fiber/v2"

// reproduce response error for swagger
type ErrorResponseSwagger struct {
	Status  string      `json:"status"`
    Message interface{} `json:"message"`
    Data    interface{} `json:"data"`
}

// global response
func JSONResponse(c *fiber.Ctx,statusCode int, status string, message interface{}, data interface{}) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

func JSONResponseWithPagination(c *fiber.Ctx, statusCode int, status string, message string, data interface{}, page int, limit int, totalRecord int64) error {
	total_pages :=(totalRecord + int64(page) -1)/ int64(page)
	return c.Status(statusCode).JSON(fiber.Map{
		"status":  status,
		"message": message,
		"data":    data,
		"paginate": fiber.Map{
			"current_page":page,
			"page":limit,
			"total":totalRecord,
			"total_page":total_pages,
		},
	})
}