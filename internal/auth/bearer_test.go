package auth_test

import (
	"net/http"
	"testing"

	"github.com/dUPYeYE/go-htmx/internal/auth"
)

func TestGetBearerToken(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		expectedToken string
		expectError   bool
	}{
		{
			name: "Valid Bearer Token",
			headers: http.Header{
				"Authorization": []string{"Bearer test-token"},
			},
			expectedToken: "test-token",
			expectError:   false,
		},
		// ... other test cases ...
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := auth.GetBearerToken(tt.headers)
			if (err != nil) != tt.expectError {
				t.Errorf("GetBearerToken() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if token != tt.expectedToken {
				t.Errorf("GetBearerToken() = %v, want %v", token, tt.expectedToken)
			}
		})
	}
}
