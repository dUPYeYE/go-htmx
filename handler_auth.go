package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/dUPYeYE/go-htmx/internal/auth"
	"github.com/dUPYeYE/go-htmx/internal/database"
)

func (cfg *config) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type requestParams struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type response struct {
		User
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}

	decoder := json.NewDecoder(r.Body)
	params := requestParams{}
	if err := decoder.Decode(&params); err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode request body")
		return
	}

	user, err := cfg.db.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Error while getting user")
		return
	}

	if err = auth.CheckPasswordHash(params.Password, user.Password); err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid password")
		return
	}

	userID, err := uuid.Parse(user.ID)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Error while parsing user ID")
		return
	}

	jwtToken, err := auth.GenerateJWT(userID, cfg.jwtSecret, time.Hour)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Error while generating JWT token")
		return
	}

	refreshToken, err := auth.GenerateRefreshToken()
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Error while generating refresh token")
		return
	}

	if _, err := cfg.db.AddRefreshToken(r.Context(), database.AddRefreshTokenParams{
		UserID: user.ID,
		Token:  refreshToken,
	}); err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Error while creating refresh token")
		return
	}

	userResp, err := databaseUserToUser(user)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Error while converting user")
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		User:         userResp,
		Token:        jwtToken,
		RefreshToken: refreshToken,
	})
}

func (cfg *config) handlerRefreshToken(w http.ResponseWriter, r *http.Request, user database.User) {
	type response struct {
		Token string `json:"token"`
	}

	bearerToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusUnauthorized, "Invalid token")
		return
	}
	refreshToken, err := cfg.db.GetRefreshToken(r.Context(), bearerToken)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Error while getting refresh token")
		return
	}
	if refreshToken.RevokedAt.Valid {
		log.Println(err)
		respondWithError(w, http.StatusUnauthorized, "Refresh token revoked")
		return
	}
	if user.ID != refreshToken.UserID {
		log.Println(err)
		respondWithError(w, http.StatusUnauthorized, "Invalid token")
		return
	}

	userID, err := uuid.Parse(refreshToken.UserID)
	jwtToken, err := auth.GenerateJWT(userID, cfg.jwtSecret, time.Hour)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Error while generating JWT token")
		return
	}

	respondWithJSON(w, http.StatusOK, response{Token: jwtToken})
}

func (cfg *config) handlerRevokeToken(w http.ResponseWriter, r *http.Request, user database.User) {
	bearerToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusUnauthorized, "Invalid token")
		return
	}

	refreshToken, err := cfg.db.GetRefreshToken(r.Context(), bearerToken)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Error while getting refresh token")
		return
	}
	if user.ID != refreshToken.UserID {
		log.Println(err)
		respondWithError(w, http.StatusUnauthorized, "Invalid token")
		return
	}

	if err = cfg.db.RevokeRefreshToken(r.Context(), database.RevokeRefreshTokenParams{
		Token:     bearerToken,
		RevokedAt: sql.NullString{String: time.Now().Format(time.RFC3339), Valid: true},
		UpdatedAt: time.Now().Format(time.RFC3339),
	}); err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Error while revoking token")
		return
	}
}
