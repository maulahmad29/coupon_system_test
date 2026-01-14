package config

import (
	"coupon_system_test/sys"
	"fmt"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog"
)

/*
Setup App
*/

func NewApp(cfg *sys.SysEnv, logger *zerolog.Logger) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: cfg.AppName,
	})

	app.Use(
		fiberzerolog.New(fiberzerolog.Config{
			Logger: logger,
		}),
		recover.New(recover.Config{
			EnableStackTrace: true,
		}),
	)

	if cfg.AppCors {
		app.Use(cors.New())
		fmt.Println("Enable cors")
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
