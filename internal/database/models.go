// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
)

type RefreshToken struct {
	Token     string
	UserID    string
	CreatedAt string
	UpdatedAt string
	ExpiresAt string
	RevokedAt sql.NullString
}

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	CreatedAt string
	UpdatedAt string
}
