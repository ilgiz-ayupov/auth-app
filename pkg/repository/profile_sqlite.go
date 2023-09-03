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

func (repo *ProfileSqlite) AddPhoneNumber(phoneNumber auth.PhoneNumber) (int64, error) {
	ctx := context.Background()

	db, err := repo.db.Conn(ctx)
	if err != nil {
		log.Fatalf("error connecting to database: %s", err.Error())
	}

	query := fmt.Sprintf("INSERT INTO %s (user_id, phone, description, is_fax) VALUES ($1, $2, $3, $4)", phoneNumbersTable)
	result, err := db.ExecContext(ctx, query, phoneNumber.UserId, phoneNumber.Phone, phoneNumber.Description, phoneNumber.IsFax)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
