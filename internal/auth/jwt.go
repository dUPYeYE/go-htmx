package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// GenerateJWT generates a new JWT token with the given userID, tokenSecret and expiresIn.
//
// Parameters:
// - userID: the user ID to be included in the token
// - tokenSecret: the secret key to sign the token
// - expiresIn: the duration of the token's validity
//
// Returns:
// - string: the generated token
// - error: an error if something goes wrong
func GenerateJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "dupp",
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(expiresIn)),
		Subject:   userID.String(),
	})

	return token.SignedString([]byte(tokenSecret))
}

// ValidateJWT validates the given JWT token with the tokenSecret.
//
// Parameters:
// - tokenString: the token to be validated
// - tokenSecret: the secret key to validate the token
//
// Returns:
// - uuid.UUID: the user ID from the token
// - error: an error if something goes wrong
func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&jwt.RegisteredClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(tokenSecret), nil
		},
	)
	if err != nil {
		return uuid.Nil, err
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if ok && token.Valid {
		return uuid.Parse(claims.Subject)
	}

	return uuid.Nil, errors.New("invalid token")
}
