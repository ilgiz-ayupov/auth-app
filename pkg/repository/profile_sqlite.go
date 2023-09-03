package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/ilgiz-ayupov/auth-app"
)

type ProfileSqlite struct {
	db *sql.DB
}

func NewProfileSqlite(db *sql.DB) *ProfileSqlite {
	return &ProfileSqlite{
		db: db,
	}
}

func (repo *ProfileSqlite) GetUser(name string) (auth.UserProfile, error) {
	log.Println(name)

	ctx := context.Background()
	var user auth.UserProfile

	db, err := repo.db.Conn(ctx)
	if err != nil {
		log.Fatalf("error connecting to database: %s", err.Error())
	}

	query := fmt.Sprintf("SELECT id, name, age FROM %s WHERE name=$1", usersTable)
	rows, err := db.QueryContext(ctx, query, name)
	if err != nil {
		log.Fatalf("error getting data from database: %s", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Name, &user.Age); err != nil {
			return auth.UserProfile{}, err
		}
	}

	return user, nil
}
