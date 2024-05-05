package auth

import (
	"errors"
	"net/http"
	"testing"
)

// TestGetAPIKey tests various scenarios for GetAPIKey function
func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		want          string
		expectedError error
	}{
		{
			name:          "No Authorization Header",
			headers:       http.Header{},
			want:          "",
			expectedError: ErrNoAuthHeaderIncluded,
		},
		{
			name: "Malformed Authorization Header - Missing Token",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			want:          "",
			expectedError: errors.New("malformed authorization header"),
		},
		{
			name: "Malformed Authorization Header - Incorrect Prefix",
			headers: http.Header{
				"Authorization": []string{"Bearer someapikey"},
			},
			want:          "",
			expectedError: errors.New("malformed authorization header"),
		},
		{
			name: "Correct Authorization Header",
			headers: http.Header{
				"Authorization": []string{"ApiKey 12345"},
			},
			want:          "12345",
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if err != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("GetAPIKey() error = %v, expectedError %v", err, tt.expectedError)
				return
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
