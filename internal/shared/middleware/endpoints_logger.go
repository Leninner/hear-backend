package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func EndpointsLogger(app *fiber.App) {
	log.Println("Available endpoints:")
	for _, route := range app.GetRoutes() {
		if route.Path != "/" || route.Method == "GET" {
			log.Printf("%s %s", route.Method, route.Path)
		}
	}
} 