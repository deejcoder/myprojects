package api

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/deejcoder/myprojects/config"
	"github.com/dgrijalva/jwt-go"
)

// ValidateAuthorization validates a JWT token in the authorization header for a request
func ValidateAuthorization(w http.ResponseWriter, r *http.Request) bool {
	auth := r.Header.Get("authorization")

	// token normally in format of Bearer {token}
	re := regexp.MustCompile("Bearer (.*)")
	match := re.FindStringSubmatch(auth)

	// test for result
	if len(match) <= 1 {
		return false
	}

	tokenString := match[1]
	// assure correct signing method
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(config.GetConfig().JwtSecret), nil
	})

	// assure token is valid
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true
	}
	return false
}
