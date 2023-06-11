package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserId   int    `json:"userId"`
	Username string `json:"username"`
	jwt.StandardClaims
}
type JWTOutput struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}
