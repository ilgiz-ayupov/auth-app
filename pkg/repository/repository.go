package repository

import (
	"database/sql"

	"github.com/ilgiz-ayupov/auth-app"
)

type Authorization interface {
	CreateUser(user auth.User) (int64, error)
	AuthentificationUser(login string, password string) (int, error)
}

type Repository struct {
	Authorization
}

func InitRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthSqlite(db),
	}
}
