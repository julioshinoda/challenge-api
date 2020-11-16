package auth

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
)

func GenerateToken(id int64) string {
	tokenAuth := jwtauth.New("HS256", []byte(os.Getenv("JWT_SECRET")), nil)
	_, tokenString, _ := tokenAuth.Encode(jwt.MapClaims{"user_id": id})
	return tokenString
}
