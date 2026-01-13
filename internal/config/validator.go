package config

import (
	"coupon_system_test/sys"

	"github.com/go-playground/validator/v10"
)

/*
Use validator to add validation tools
*/

func NewValidator(cfg *sys.SysEnv) *validator.Validate {
	return validator.New()
}
