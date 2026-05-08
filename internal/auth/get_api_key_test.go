package auth

import (
	"net/http"
	"testing"
)

func TestCorrectGetAPIKEY(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "ApiKey some-secret-key")
	key, err := GetAPIKey(headers)
	if key != "some-secret-key" {
		t.Fatalf("Expected %v, got %v", "some-secret-key", key)
	}
	newHeaders := http.Header{}
	_, err = GetAPIKey(newHeaders)
	if err == nil {
		t.Fatalf("Expected %v, got %v", "no authorization header included", err)
	}
}
