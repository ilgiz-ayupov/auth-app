package repository

import (
	"database/sql"

	"github.com/ilgiz-ayupov/auth-app"
)

type Authorization interface {
	CreateUser(user auth.User) (int64, error)
	AuthentificationUser(login string, password string) (int, error)
	AuthorizationToken(claims auth.UserTokenClaims) error
}

type Profile interface {
	GetUser(name string) (auth.UserProfile, error)
}

type Repository struct {
	Authorization
	Profile
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthSqlite(db),
		Profile:       NewProfileSqlite(db),
	}
}
