package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))

type Claims struct {
	UserID       string `json:"user_id"`
	TokenVersion int    `json:"token_version"`
	IsAdmin      bool   `json:"is_admin"`
	jwt.RegisteredClaims
}

func GenerateToken(UserID string, TokenVersion int, IsAdmin bool) (string, error) {
	claims := Claims{
		UserID:       UserID,
		TokenVersion: TokenVersion,
		IsAdmin:      IsAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
