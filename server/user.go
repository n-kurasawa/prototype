package main

import "time"

type Repository interface {
	getUser() *User
}

type User struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Age       int       `db:"age"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
