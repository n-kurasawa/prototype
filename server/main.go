package main

import (
	"flag"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	port       = flag.Int("port", 8080, "server port")
	configFile = flag.String("config", "config.yaml", "app server configuration file")
)

func main() {
	flag.Parse()

	cfg, err := loadFile(*configFile)
	if err != nil {
		panic(err)
	}

	db, err := sqlx.Open("mysql", cfg.MySQL.DataSource())
	if err != nil {
		panic(err)
	}
	repo := NewRepository(db)
	handler := NewHandler(repo)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	e.GET("/", handler.userHello)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", *port)))
}
