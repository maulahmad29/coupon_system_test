package http

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Route struct {
	App *fiber.App
}

func NewRoute(app *fiber.App, session_start_at time.Time) Route {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(
			map[string]interface{}{
				"success": "ok",
			},
		)
	})

	app.Get("/health-check", func(c *fiber.Ctx) error {

		return c.Status(fiber.StatusOK).JSON(
			map[string]interface{}{
				"session_start_at": session_start_at,
				"time_server":      time.Now(),
			},
		)
	})

	return Route{
		App: app,
	}
}
