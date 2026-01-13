package config

import (
	"coupon_system_test/sys"
	"os"

	"github.com/rs/zerolog"
)

/*
Use zerlog for loggin app access
*/

func NewZerolog(cfg *sys.SysEnv) zerolog.Logger {

	log := zerolog.New(os.Stderr).Level(zerolog.TraceLevel).With().Timestamp().Logger()

	return log

}
