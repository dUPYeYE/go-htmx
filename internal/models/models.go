package models

import (
	"time"

	"github.com/google/uuid"

	"github.com/dUPYeYE/go-htmx/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// databaseUserToUser converts a database.User to a User.
// It returns an error if the conversion fails
//
// Parameters:
// dbUser: the database.User to convert
//
// Returns:
// User: the converted User
// error: an error if the conversion fails
func DatabaseUserToUser(dbUser database.User) (User, error) {
	return User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		Email:     dbUser.Email,
		Password:  dbUser.Password,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}, nil
}
