package auth

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateRefreshToken generates a random 32-byte token
// and returns it as a hex-encoded string.
//
// Returns:
// - string: the generated token
// - error: an error if the random number generator fails
func GenerateRefreshToken() (string, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(token), nil
}
