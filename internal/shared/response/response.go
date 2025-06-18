package response

import "github.com/gofiber/fiber/v2"

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  []string    `json:"errors,omitempty"`
}

func Success(c *fiber.Ctx, message string, data interface{}) error {
	return c.JSON(Response{
		Message: message,
		Data:    data,
	})
}

func Error(c *fiber.Ctx, status int, message interface{}) error {
	resp := Response{}

	switch v := message.(type) {
	case []string:
		resp.Errors = v
		resp.Message = "Validation failed"
	default:
		resp.Message = v.(string)
	}

	return c.Status(status).JSON(resp)
} 