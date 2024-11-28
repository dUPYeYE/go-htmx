package auth_test

import (
	"encoding/hex"
	"testing"

	"github.com/dUPYeYE/go-htmx/internal/auth"
)

func TestGenerateRefreshToken(t *testing.T) {
	token, err := auth.GenerateRefreshToken()
	if err != nil {
		t.Fatalf("GenerateRefreshToken() error = %v", err)
	}

	// Verify token length (32 bytes = 64 hex characters)
	if len(token) != 64 {
		t.Errorf("GenerateRefreshToken() token length = %v, want 64", len(token))
	}

	// Verify token is valid hex
	_, err = hex.DecodeString(token)
	if err != nil {
		t.Errorf("GenerateRefreshToken() generated invalid hex string: %v", err)
	}
}
