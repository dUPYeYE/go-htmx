package main

import (
	"net/http"

	"github.com/dUPYeYE/go-htmx/internal/auth"
	"github.com/dUPYeYE/go-htmx/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

// middlewareAuth is a middleware that checks if the request is authenticated.
// If the request is authenticated, it calls the handler with the authenticated user.
//
// Parameters:
// - handler: authedHandler - The handler to be called if the request is authenticated
//
// Returns:
// - http.HandlerFunc: The handler function
func (cfg *config) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jwtToken, err := auth.GetBearerToken(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "No token found")
			return
		}

		userID, err := auth.ValidateJWT(jwtToken, cfg.jwtSecret)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		user, err := cfg.db.GetUserById(r.Context(), userID)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Error while fetching user")
			return
		}

		handler(w, r, user)
	}
}
