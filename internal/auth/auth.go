package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts API Key
// from http request headers
// Example :
// Authorization: "ApiKey {APIKEY}"
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if len(val) < 8 {
		return "", errors.New("bad authorization header")
	}
	if strings.ToLower(val[0:7]) != "apikey " {
		return "", errors.New("bad authorization header")
	}
	return val[7:], nil
}
