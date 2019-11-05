package helpers

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var JWTSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(id string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	t, _ := token.SignedString(JWTSecret)
	return t
}
