package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetBearerToken extracts the Bearer token from the Authorization header.
//
// Parameters:
// - headers: the headers from the HTTP request
//
// Returns:
// - string: the Bearer token
// - error: an error if something goes wrong
func GetBearerToken(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("missing Authorization header")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", errors.New("invalid Authorization header")
	}

	return parts[1], nil
}
