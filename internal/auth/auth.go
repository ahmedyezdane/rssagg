package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Example:
// Authorization: ApiKey {api_key}
func GetAPIKey(headers http.Header) (string, error) {
	header := headers.Get("Authorization")

	if header == "" {
		return "", errors.New("No authentication info found")
	}

	values := strings.Split(header, " ")

	if len(values) != 2 || values[0] != "ApiKey" {
		return "", errors.New("Malformed auth header")
	}

	return values[1], nil
}
