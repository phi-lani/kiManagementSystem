package utils

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// Claims structure
type Claims struct {
	UserID   uint   `json:"userID"`   // Add userID to the claims
	Username string `json:"username"` // Keep the username in case you need it
	Role     string `json:"role"`
	jwt.StandardClaims
}

// GenerateJWT creates a new token with userID, username, and role
func GenerateJWT(userID uint, username string, role string) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24) // Token expires in 24 hours
	claims := &Claims{
		UserID:   userID,
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Create the token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func VerifyJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
