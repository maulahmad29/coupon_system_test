package main

import (
	"coupon_system_test/db/migration"
	"coupon_system_test/internal/config"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2/log"
)

func main() {
	log.Info("Main start load")
	initViper := config.NewViper()
	initZeroLogger := config.NewZerolog(initViper)
	initDbSqlx := config.NewSqlx(initViper)
	initValidate := config.NewValidator(initViper)
	if initViper.AppEnv != "production" {
		migration.RunMigration(initViper)
	}

	initApp := config.NewApp(initViper, &initZeroLogger)

	config.NewBootstrap(&config.Boostrap{
		Cfg:      initViper,
		Logger:   &initZeroLogger,
		Validate: initValidate,
		App:      initApp,
		DbSqlx:   initDbSqlx,
	}, time.Now())

	go func() {
		err := initApp.Listen(fmt.Sprintf(":%s", initViper.AppPort))
		if err != nil {
			if initViper.AppEnv != "production" {
				migration.DropDatabase(initDbSqlx)
			}
			initDbSqlx.Close()
			log.Fatal("Failed to start server :", err)
		}

	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	<-quit
	if initViper.AppEnv != "production" {
		migration.DropDatabase(initDbSqlx)
	}
	initDbSqlx.Close()
	log.Info("Connection has been close")
	log.Info("Fiber has been shutdown")
	initApp.Shutdown()
	log.Info("Shutdown complete...")
	os.Exit(0)

}
