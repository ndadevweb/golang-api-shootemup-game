package jwt

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("fdlkfdscqj,dksr,c'é!çc(n'!ç(nç!n'c(DSgfHGL!ç('QSDfghjT:,:;)")

func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(15 * time.Minute).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}
