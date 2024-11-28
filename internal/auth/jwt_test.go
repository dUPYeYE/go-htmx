package auth_test

import (
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/dUPYeYE/go-htmx/internal/auth"
)

func TestGenerateAndValidateJWT(t *testing.T) {
	userID := uuid.New()
	tokenSecret := "test-secret"
	expiresIn := time.Hour

	// Test token generation
	token, err := auth.GenerateJWT(userID, tokenSecret, expiresIn)
	if err != nil {
		t.Fatalf("GenerateJWT() error = %v", err)
	}
	if token == "" {
		t.Error("GenerateJWT() returned empty token")
	}

	// Test token validation
	parsedUserID, err := auth.ValidateJWT(token, tokenSecret)
	if err != nil {
		t.Fatalf("ValidateJWT() error = %v", err)
	}
	if parsedUserID != userID {
		t.Errorf("ValidateJWT() userID = %v, want %v", parsedUserID, userID)
	}

	// Test invalid token
	_, err = auth.ValidateJWT("invalid-token", tokenSecret)
	if err == nil {
		t.Error("ValidateJWT() expected error for invalid token")
	}

	// Test wrong secret
	_, err = auth.ValidateJWT(token, "wrong-secret")
	if err == nil {
		t.Error("ValidateJWT() expected error for wrong secret")
	}
}
