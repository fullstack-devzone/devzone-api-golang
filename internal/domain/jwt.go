package domain

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/config"
)

type Claims struct {
	jwt.RegisteredClaims
	UserId   int    `json:"userId"`
	Username string `json:"username"`
}
type JWTOutput struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}

func CreateJwtToken(cfg config.AppConfig, user User) (JWTOutput, error) {
	expirationTime := time.Now().Add(2 * time.Hour)
	claims := Claims{
		UserId:   user.Id,
		Username: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cfg.JwtSecret))
	if err != nil {
		return JWTOutput{}, err
	}
	jwtOutput := JWTOutput{
		Token:   tokenString,
		Expires: expirationTime,
	}
	return jwtOutput, nil
}
