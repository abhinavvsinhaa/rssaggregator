package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Authorization: ApiKey {insert apiKey here}
func ExtractAPIKeyFromHeader(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("No authentication info found")
	}

	split := strings.Split(val, " ")
	if len(split) != 2 || split[0] != "ApiKey" {
		return "", errors.New("Malformed auth header")
	}

	return split[1], nil
}
