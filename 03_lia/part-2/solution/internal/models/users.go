package models

import (
	"database/sql"
	"fmt"
)

// User holds data about a user.
type User struct {
	ID       int
	Username string
	Email    string
}

// UserModel holds a database handle to manipulate a User.
type UserModel struct {
	DB *sql.DB
}

// NewUserModel creates a Users table and returns a new UserModel.
func NewUserModel(db *sql.DB) (*UserModel, error) {
	m := UserModel{db}
	err := m.createTable()
	if err != nil {
		return nil, fmt.Errorf("creating table: %w", err)
	}
	return &m, nil
}

// createTable creates a Users table.
func (m *UserModel) createTable() error {
	stmt := `
		CREATE TABLE IF NOT EXISTS Users (
		    id INTEGER PRIMARY KEY,
		    username TEXT NOT NULL,
		    email TEXT NOT NULL UNIQUE,
		    password TEXT NOT NULL
		);
	`
	_, err := m.DB.Exec(stmt)
	if err != nil {
		return fmt.Errorf("creating Users table: %w", err)
	}
	return nil
}
