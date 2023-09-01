package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/ilgiz-ayupov/auth-app"

	_ "github.com/mattn/go-sqlite3"
)

type AuthSqlite struct {
	db *sql.DB
}

func NewAuthSqlite(db *sql.DB) *AuthSqlite {
	return &AuthSqlite{
		db: db,
	}
}

func (repo *AuthSqlite) CreateUser(user auth.User) (string, error) {
	ctx := context.Background()

	db, err := repo.db.Conn(ctx)
	if err != nil {
		log.Fatalf("error connecting to database: %s", err.Error())
	}

	query := fmt.Sprintf("INSERT INTO users (login, password, name, age) VALUES ($1, $2, $3, $4)")
	_, err = db.ExecContext(ctx, query, user.Login, user.Password, user.Name, user.Age)
	if err != nil {
		return "", err
	}

	return user.Name, nil
}
