package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey_Success(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey secret123")

	key, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if key != "secret123" {
		t.Errorf("expected key %q, got %q", "secret123", key)
	}
}

func TestGetAPIKey_NoAuthHeader(t *testing.T) {
	headers := http.Header{} // no Authorization header

	_, err := GetAPIKey(headers)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err != ErrNoAuthHeaderIncluded {
		t.Errorf("expected %v, got %v", ErrNoAuthHeaderIncluded, err)
	}
}
