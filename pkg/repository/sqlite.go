package repository

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/sqlite3"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/mattn/go-sqlite3"
)

const (
	usersTable = "users"
)

type Config struct {
	Path string
}

func OpenSqliteDB(cfg *Config) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", cfg.Path)
	if err != nil {
		return nil, err
	}

	if err := acceptMigrations(db); err != nil {
		return nil, err
	}

	return db, nil
}

func acceptMigrations(db *sql.DB) error {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./db/migrations",
		"sqlite3", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	version, dirty, err := m.Version()
	if err != nil {
		return err
	}

	fmt.Printf("Current version migrations: %d, dirty: %t\n", version, dirty)
	return nil
}
