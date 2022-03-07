package main

import "github.com/jmoiron/sqlx"

type RepositoryImpl struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *RepositoryImpl {
	return &RepositoryImpl{db: db}
}
func (r *RepositoryImpl) getUser() *User {
	var user User
	if err := r.db.Get(&user, "select * from users"); err != nil {
		panic(err)
	}
	return &user
}
