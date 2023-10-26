package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// DB represents the PostgreSQL database.
type DB struct {
	*sql.DB
}

// NewDB creates a new database connection.
func NewDB(host, port, user, password, dbName string) (*DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

// Close closes the database connection.
func (db *DB) Close() {
	db.DB.Close()
}
