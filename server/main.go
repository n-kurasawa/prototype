package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db, err := sqlx.Open("mysql", "root:@tcp(db:3306)/prototype?parseTime=true")
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
	e.Logger.Fatal(e.Start(":8080"))
}
