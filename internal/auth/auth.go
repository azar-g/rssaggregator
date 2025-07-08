package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")

	if authHeader == "" {

		return "", errors.New("authorization header is missing")
	}

	authHeaderParts := strings.Split(authHeader, " ")

	if len(authHeaderParts) != 2 {
		return "", errors.New("authorization header is malformed")
	}

	if authHeaderParts[0] != "ApiKey" {
		return "", errors.New("authorization header first part is malformed")
	}

	return authHeaderParts[1], nil
}
