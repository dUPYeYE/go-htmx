package auth_test

import (
	"testing"

	"github.com/dUPYeYE/go-htmx/internal/auth"
)

func TestHashPasswordAndCheck(t *testing.T) {
	password := "test-password"

	// Test password hashing
	hash, err := auth.HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword() error = %v", err)
	}
	if hash == "" {
		t.Error("HashPassword() returned empty hash")
	}
	if hash == password {
		t.Error("HashPassword() returned unhashed password")
	}

	// Test correct password check
	err = auth.CheckPasswordHash(password, hash)
	if err != nil {
		t.Errorf("CheckPasswordHash() error = %v", err)
	}

	// Test incorrect password
	err = auth.CheckPasswordHash("wrong-password", hash)
	if err == nil {
		t.Error("CheckPasswordHash() expected error for wrong password")
	}
}
