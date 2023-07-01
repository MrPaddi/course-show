package main

// use fiber bootstrap web serve
import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/hello/:name", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("Hello, %s 👋!", c.Params("name")))
	})

	app.Listen(":8000")
}
