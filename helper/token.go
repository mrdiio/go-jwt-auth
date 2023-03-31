package helper

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(payload interface{}, e int, secretJwtKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"payload": payload,
		"exp":     time.Now().Add(time.Second * time.Duration(e)).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretJwtKey))

	if err != nil {
		return "", fmt.Errorf("error generating token: %v", err)
	}

	return tokenString, nil
}

func ValidateToken(requestToken string, secretJwtKey string) (jwt.MapClaims, error) {
	t, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretJwtKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil

}
