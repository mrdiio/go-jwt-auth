package helper

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(ttl time.Duration, payload interface{}, secretJWTKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"payload": payload,
		"exp":     time.Now().Add(ttl).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretJWTKey))

	if err != nil {
		return "", fmt.Errorf("error generating token: %v", err)
	}

	return tokenString, nil
}

func VerifyToken(token string, signedJWTKey string) (jwt.MapClaims, error) {
	t, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", jwtToken.Header["alg"])
		}

		return []byte(signedJWTKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing token: %v", err)
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
