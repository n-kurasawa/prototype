package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Open("mysql", "root:@tcp(localhost:33060)/prototype?parseTime=true")
	if err != nil {
		panic(err)
	}
	repo := NewRepository(db)
	handler := NewHandler(repo)

	http.HandleFunc("/", handler.userHello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
