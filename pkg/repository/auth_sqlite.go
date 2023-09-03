package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/ilgiz-ayupov/auth-app"
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
	conn, ctx, err := ConnSqliteDB(repo.db)
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf("INSERT INTO users (login, password, name, age) VALUES ($1, $2, $3, $4)")
	result, err := conn.ExecContext(ctx, query, user.Login, user.Password, user.Name, user.Age)
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
	conn, ctx, err := ConnSqliteDB(repo.db)
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf("SELECT id FROM %s WHERE login=$1 AND password=$2", usersTable)
	rows, err := conn.QueryContext(ctx, query, login, password)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var id int
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

func (repo *AuthSqlite) AuthorizationToken(claims auth.UserTokenClaims) error {
	conn, ctx, err := ConnSqliteDB(repo.db)
	if err != nil {
		return err
	}

	query := fmt.Sprintf("SELECT id FROM %s WHERE id=$1 AND login=$2", usersTable)
	rows, err := conn.QueryContext(ctx, query, claims.UserId, claims.Login)
	defer rows.Close()

	var id int
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return err
		}
	}

	if id == 0 {
		return errors.New("error user not found")
	}

	return nil
}
