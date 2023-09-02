package repository

import (
	"context"
	"database/sql"
	"errors"
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

func (repo *AuthSqlite) CreateUser(user auth.User) (int64, error) {
	ctx := context.Background()

	db, err := repo.db.Conn(ctx)
	if err != nil {
		log.Fatalf("error connecting to database: %s", err.Error())
	}

	query := fmt.Sprintf("INSERT INTO users (login, password, name, age) VALUES ($1, $2, $3, $4)")
	result, err := db.ExecContext(ctx, query, user.Login, user.Password, user.Name, user.Age)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (repo *AuthSqlite) AuthentificationUser(login string, password string) (int, error) {
	var id int
	ctx := context.Background()

	db, err := repo.db.Conn(ctx)
	if err != nil {
		log.Fatalf("error connecting to database: %s", err.Error())
	}

	query := fmt.Sprintf("SELECT id FROM %s WHERE login=$1 AND password=$2", usersTable)
	rows, err := db.QueryContext(ctx, query, login, password)
	if err != nil {
		log.Fatalf("error getting data from database: %s", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return 0, err
		}
	}

	if id == 0 {
		return 0, errors.New("error user not found")
	}

	return id, nil
}
