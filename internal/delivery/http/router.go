package http

import (
	"coupon_system_test/internal/handler"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Route struct {
	App *fiber.App
}

func NewRoute(app *fiber.App, session_start_at time.Time, couponHandler handler.ICouponHandler) Route {

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

	app.Post("api/coupons/", couponHandler.CreateCoupon)

	app.Get("api/coupons/:coupon_name", couponHandler.GetCouponDetailClaims)

	app.Post("api/coupons/claim", couponHandler.ClaimCoupon)

	return Route{
		App: app,
	}
}
