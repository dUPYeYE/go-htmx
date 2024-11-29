package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/dUPYeYE/go-htmx/internal/auth"
	"github.com/dUPYeYE/go-htmx/internal/database"
	"github.com/dUPYeYE/go-htmx/internal/models"
)

// CreateUser creates a new user with the given name, email and password.
func (cfg *config) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type requestParams struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	params := requestParams{}
	if err := decoder.Decode(&params); err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode request body")
		return
	}

	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Error while hashing password")
		return
	}

	user, err := cfg.db.CreateUser(r.Context(), database.CreateUserParams{
		Name:     params.Name,
		Email:    params.Email,
		Password: hashedPassword,
	})
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Error while creating user")
		return
	}

	respUser, err := models.DatabaseUserToUser(user)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Error while converting user")
		return
	}
	respondWithJSON(w, http.StatusCreated, respUser)
}

// GetUsers returns all the users
func (cfg *config) handlerGetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := cfg.db.GetAllUsers(r.Context())
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Error while fetching users")
		return
	}

	respUsers := make([]models.User, 0, len(users))
	for _, user := range users {
		respUser, err := models.DatabaseUserToUser(user)
		if err != nil {
			log.Println(err)
			respondWithError(w, http.StatusInternalServerError, "Error while converting user")
			return
		}
		respUsers = append(respUsers, respUser)
	}

	respondWithJSON(w, http.StatusOK, respUsers)
}

// GetUser returns the user with the given ID.
func (cfg *config) handlerGetOneUser(w http.ResponseWriter, r *http.Request, user database.User) {
	if user.ID.String() != chi.URLParam(r, "id") {
		respondWithError(w, http.StatusForbidden, "You can only get your own account")
		return
	}

	respUser, err := models.DatabaseUserToUser(user)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Error while converting user")
		return
	}

	respondWithJSON(w, http.StatusOK, respUser)
}

// DeleteUser deletes the user with the given ID.
func (cfg *config) handlerDeleteUser(w http.ResponseWriter, r *http.Request, user database.User) {
	id := chi.URLParam(r, "id")
	if id != user.ID.String() {
		respondWithError(w, http.StatusForbidden, "You can only delete your own account")
		return
	}

	if err := cfg.db.DeleteUser(r.Context(), user.ID); err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Error while deleting user")
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)
}
