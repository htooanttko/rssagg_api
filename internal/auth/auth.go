package auth

import (
	"errors"
	"net/http"
	"strings"
)

var ErrNoAuthHeader = errors.New("no authorization header included")
var ErrMalHeader = errors.New("malformed authorization header")

func GetAPIKey(headers http.Header) (string,error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeader
	} 

	keyString := strings.Split(authHeader, " ")

	if len(keyString) < 2 || keyString[0] != "ApiKey" {
		return "", ErrMalHeader
	}


	return keyString[1], nil
}