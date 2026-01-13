package config

import (
	"coupon_system_test/sys"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rs/zerolog"
)

/*
Setup App
*/

func NewApp(cfg *sys.SysEnv, logger *zerolog.Logger) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: cfg.AppName,
	})

	if cfg.AppCors {
		app.Use(cors.New())
	}

	return app
}

func NewDefaultErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		return ctx.Status(code).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
}
