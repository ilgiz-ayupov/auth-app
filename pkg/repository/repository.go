package repository

import (
	"database/sql"

	"github.com/ilgiz-ayupov/auth-app"
)

type Authorization interface {
	CreateUser(user auth.User) (string, error)
}

type Repository struct {
	Authorization
}

func InitRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthSqlite(db),
	}
}
