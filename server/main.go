package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Open("mysql", "root:password@tcp(localhost:3306)/prottype")
	if err != nil {
		panic(err)
	}
	repo := NewRepository(db)
	handler := NewHandler(repo)

	http.HandleFunc("/", handler.userHello)
	http.ListenAndServe(":8080", nil)
}
