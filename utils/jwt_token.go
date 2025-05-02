package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var tokenSecret string

func New(ts string) {

	tokenSecret = ts
}

func GenerateToken(id string, identifier string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"userId": id, "identifier": identifier, "exp": time.Now().Add(2 * time.Hour).Unix()})
	return token.SignedString([]byte(tokenSecret))
}
func ValidateToken(token string) (string, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return "", err
	}
	var id string
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
		id = claims["userId"].(string)
	}
	return id, nil
}
