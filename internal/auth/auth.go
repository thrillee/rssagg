package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts an API Key from
// the headers of an HTTP requests
// Example
// Authorization: Bearer {API TOKEN}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")

	if val == "" {
		return "", errors.New("No authentication info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("Malformed auth header")
	}

	bearer := vals[0]
	if bearer != "Bearer" {
		return "", errors.New("Unknown auth bearer")
	}

	apiKey := vals[1]
	return apiKey, nil
}
