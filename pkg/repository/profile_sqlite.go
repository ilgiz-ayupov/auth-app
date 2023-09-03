package repository

import (
	"database/sql"
	"errors"
	"fmt"

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
	conn, ctx, err := ConnSqliteDB(repo.db)
	if err != nil {
		return auth.UserProfile{}, err
	}

	query := fmt.Sprintf("SELECT id, name, age FROM %s WHERE name=$1", usersTable)
	rows, err := conn.QueryContext(ctx, query, name)
	if err != nil {
		return auth.UserProfile{}, err
	}
	defer rows.Close()

	var user auth.UserProfile
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Name, &user.Age); err != nil {
			return auth.UserProfile{}, err
		}
	}

	return user, nil
}

func (repo *ProfileSqlite) SearchPhoneNumbers(phone string) ([]auth.PhoneNumber, error) {
	conn, ctx, err := ConnSqliteDB(repo.db)
	if err != nil {
		return []auth.PhoneNumber{}, err
	}

	query := fmt.Sprintf("SELECT id, user_id, phone, description, is_fax FROM %s WHERE phone LIKE $1", phoneNumbersTable)
	rows, err := conn.QueryContext(ctx, query, "%"+phone+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var phoneNums []auth.PhoneNumber
	for rows.Next() {

		var phoneNum auth.PhoneNumber
		if err := rows.Scan(&phoneNum.Id, &phoneNum.UserId, &phoneNum.Phone, &phoneNum.Description, &phoneNum.IsFax); err != nil {
			return []auth.PhoneNumber{}, err
		}

		phoneNums = append(phoneNums, phoneNum)
	}

	if len(phoneNums) == 0 {
		return []auth.PhoneNumber{}, errors.New("phone number is not found")
	}
	return phoneNums, nil
}

func (repo *ProfileSqlite) AddPhoneNumber(phoneNum auth.PhoneNumber) (int64, error) {
	conn, ctx, err := ConnSqliteDB(repo.db)
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf("INSERT INTO %s (user_id, phone, description, is_fax) VALUES ($1, $2, $3, $4)", phoneNumbersTable)
	result, err := conn.ExecContext(ctx, query, phoneNum.UserId, phoneNum.Phone, phoneNum.Description, phoneNum.IsFax)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (repo *ProfileSqlite) UpdatePhoneNumber(updating auth.UpdatingPhoneNumber) error {
	conn, ctx, err := ConnSqliteDB(repo.db)
	if err != nil {
		return err
	}

	query := fmt.Sprintf("SELECT id, user_id, phone, description, is_fax FROM %s WHERE id=$1", phoneNumbersTable)
	rows, err := conn.QueryContext(ctx, query, updating.Id)
	if err != nil {
		return err
	}
	defer rows.Close()

	var phoneNum auth.PhoneNumber
	for rows.Next() {
		if err := rows.Scan(&phoneNum.Id, &phoneNum.UserId, &phoneNum.Phone, &phoneNum.Description, &phoneNum.IsFax); err != nil {
			return err
		}
	}

	if phoneNum.UserId != updating.UserId {
		query := fmt.Sprintf("UPDATE %s SET user_id=$1 WHERE id=$2", phoneNumbersTable)

		if _, err = conn.ExecContext(ctx, query, updating.UserId, updating.Id); err != nil {
			return err
		}
	}

	if phoneNum.Phone != updating.Phone {
		query := fmt.Sprintf("UPDATE %s SET phone=$1 WHERE id=$2", phoneNumbersTable)

		if _, err = conn.ExecContext(ctx, query, updating.Phone, updating.Id); err != nil {
			return err
		}
	}

	if phoneNum.Description != updating.Description {
		query := fmt.Sprintf("UPDATE %s SET description=$1 WHERE id=$2", phoneNumbersTable)

		if _, err = conn.ExecContext(ctx, query, updating.Description, updating.Id); err != nil {
			return err
		}
	}

	if phoneNum.IsFax != updating.IsFax {
		query := fmt.Sprintf("UPDATE %s SET is_fax=$1 WHERE id=$2", phoneNumbersTable)

		if _, err = conn.ExecContext(ctx, query, updating.IsFax, updating.Id); err != nil {
			return err
		}
	}

	return nil
}

func (repo *ProfileSqlite) DeletePhoneNumber(id int) error {
	conn, ctx, err := ConnSqliteDB(repo.db)
	if err != nil {
		return err
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", phoneNumbersTable)
	_, err = conn.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
