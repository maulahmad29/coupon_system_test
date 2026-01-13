package config

import (
	"coupon_system_test/sys"
	"fmt"

	"github.com/spf13/viper"
)

/*
use viper to load env
*/

func NewViper() *sys.SysEnv {
	fmt.Println("Start load config with viper...")
	viperInstance := viper.New()
	viperInstance.SetConfigFile(".env")
	viperInstance.AutomaticEnv()

	if err := viperInstance.ReadInConfig(); err != nil {
		fmt.Println("Failed read config viper to .env ....")
		panic(err)
	}

	cfg := new(sys.SysEnv)
	if err := viperInstance.Unmarshal(cfg); err != nil {
		fmt.Println("Failed unamrshal config viper .....")
		panic(err)
	}

	fmt.Println("Finish load config with viper ...")

	return cfg
}
