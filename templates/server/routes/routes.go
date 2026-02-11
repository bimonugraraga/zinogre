package routes

import (
	"github.com/bimonugraraga/zinogre/templates/server/handlers"
	"github.com/gofiber/fiber/v2"
)

func InitFiber() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello Fiber",
		})
	})
	app.Get("/hello", handlers.HelloWorldHandler)
	app.Listen(":8080")
}
