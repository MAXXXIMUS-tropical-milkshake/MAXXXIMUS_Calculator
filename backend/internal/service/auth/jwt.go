package auth

import (
	"github.com/golang-jwt/jwt"
)

func getJWTToken(jwtSecret string, userID int) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
	})

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}
