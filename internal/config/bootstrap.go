package config

import (
	"coupon_system_test/internal/delivery/http"
	"coupon_system_test/internal/handler"
	"coupon_system_test/internal/repo"
	"coupon_system_test/internal/usecase"
	"coupon_system_test/sys"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

/*
Bootstrap the pattern
*/

type Boostrap struct {
	App      *fiber.App
	Cfg      *sys.SysEnv
	Logger   *zerolog.Logger
	Validate *validator.Validate
	DbSqlx   *sqlx.DB
}

func NewBootstrap(cfg *Boostrap, session_start_at time.Time) {

	couponRepo := repo.NewCouponRepo()
	couponUseCase := usecase.NewCouponUseCase(cfg.DbSqlx, couponRepo)
	couponHandler := handler.NewCouponHander(cfg.Validate, couponUseCase)
	http.NewRoute(cfg.App, session_start_at, couponHandler)
}
