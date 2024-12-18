package main

import (
	tmpl "example.com/m/templates"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return Render(c, tmpl.Home())
	})
	app.Get("/ping", func(c *fiber.Ctx) error {
		return Render(c, tmpl.Ping())
	})
	app.Get("/pong", func(c *fiber.Ctx) error {
		return Render(c, tmpl.Pong())
	})

	app.Static("/", "./static")

	app.Listen(":8000")
}

func Render(c *fiber.Ctx, component templ.Component) error {
	headers := c.GetReqHeaders()

	cmp := component
	if _, isHTMXRequest := headers["HX-Request"]; !isHTMXRequest {
		cmp = tmpl.Layout(component)
	}

	c.Set("Content-Type", "text/html")
	return cmp.Render(c.Context(), c.Response().BodyWriter())
}
