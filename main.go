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

	home := tmpl.Home()
	ping := tmpl.SendBtn("ping")
	pong := tmpl.SendBtn("pong")

	app.Get("/", func(c *fiber.Ctx) error {
		return Render(c, home)
	})
	app.Get("/ping", func(c *fiber.Ctx) error {
		return Render(c, pong)
	})
	app.Get("/pong", func(c *fiber.Ctx) error {
		return Render(c, ping)
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
