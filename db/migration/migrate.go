package migration

import (
	"coupon_system_test/sys"
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

func RunMigration(cfg *sys.SysEnv) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DbUSer,
		cfg.DbPassword,
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbName,
	)
	m, err := migrate.New(
		"file:///app/db/migration",
		dsn,
	)

	if err != nil {
		log.Fatalf("migration init failed: %v", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migration failed: %v", err)
	}

	log.Info("database migration completed")
}

func DropDatabase(DbSqlx *sqlx.DB) {
	_, err := DbSqlx.Exec(`
		DROP SCHEMA IF EXISTS public CASCADE;
		CREATE SCHEMA public;
	`)
	if err != nil {
		log.Info("failed to drop schema: %v", err)
		return
	}

	log.Info("database schema dropped successfully")
}
