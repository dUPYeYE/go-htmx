package auth

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes the given password.
//
// Parameters:
// - password: the password to be hashed
//
// Returns:
// - string: the hashed password
// - error: an error if something goes wrong
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CheckPasswordHash compares the given password with the hash.
//
// Parameters:
// - password: the password to be checked
// - hash: the hash to be compared with
//
// Returns:
// - error: an error if the password and hash do not match
func CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
