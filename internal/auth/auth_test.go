package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name      string
		headers   http.Header
		key       string
		expectErr bool
	}{
		{
			name:      "valid API key",
			headers:   http.Header{"Authorization": []string{"ApiKey my-apikey"}},
			key:       "my-apikey",
			expectErr: false,
		},
		{
			name:      "invalid API key",
			headers:   http.Header{"Authorization": []string{"Token my-apikey"}},
			key:       "",
			expectErr: true,
		},
		{
			name:      "Invalid authorization header",
			headers:   http.Header{},
			key:       "",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got error: %v", err, tt.expectErr)
			}
			if gotKey != tt.key {
				t.Errorf("expected key: %s, got key: %s", tt.key, gotKey)
			}
		})
	}
}
