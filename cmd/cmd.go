package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/nomionz/currency-converter/internal/api"
)

func RunApi() error {
	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	api.NewServer().Router(app)

	return app.Listen(":80")
}
