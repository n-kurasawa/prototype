package main

type Repository interface {
	getUser() *User
}

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}
