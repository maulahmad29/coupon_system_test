package config

import (
	"coupon_system_test/sys"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

/*
Use sqlx as query builder from postgres
*/

func NewSqlx(cfg *sys.SysEnv) *sqlx.DB {
	fmt.Println("Start load config database ....")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbUSer,
		cfg.DbPassword,
		cfg.DbName,
	)

	dbConn, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		fmt.Println("Failed open connection database ....")
		panic(err)
	}

	dbConn.SetMaxOpenConns(int(cfg.DbMaxOpenConn))
	dbConn.SetConnMaxLifetime(time.Duration(cfg.DbMaxConnLifeTime) * time.Second)

	fmt.Println("Finish load config database ...")

	return dbConn
}
