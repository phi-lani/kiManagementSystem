package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// Updated Claims structure to include userID
type Claims struct {
	UserID   uint   `json:"userId"`   // Add userID to the claims
	Username string `json:"username"` // Keep the username in case you need it
	jwt.StandardClaims
}

// GenerateJWT creates a new token with userID and username
func GenerateJWT(userID uint, username string) (string, error) {
	expirationTime := time.Now().Add(time.Hour)
	claims := &Claims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// VerifyJWT verifies the token and returns the claims
func VerifyJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
