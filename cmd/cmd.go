package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/nomionz/currency-converter/internal/api"
)

func RunApi() error {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	api.NewServer().Router(app)

	return app.Listen(":80")
}
