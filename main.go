package main

// use fiber bootstrap web serve
import (
	"flag"
	"fmt"
	"log"

	"github.com/MrPaddi/course-show/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// get configFile from cmd args --config
	configFile := flag.String("config", "config.json", "Configuration file in json format")
	flag.Parse()

	// init config
	conf := config.New(*configFile)
	log.Println(conf)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/hello/:name", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("Hello, %s 👋!", c.Params("name")))
	})

	app.Listen(":" + conf.Port)
}
