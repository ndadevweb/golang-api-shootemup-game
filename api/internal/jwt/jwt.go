package jwt

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
	Config "gowebgame/internal/config"
)


func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": username,
		"exp":      time.Now().Add(15 * time.Minute).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(Config.JwtKey)
}
