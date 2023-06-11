package domain

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/config"
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

func CreateJwtToken(cfg config.AppConfig, user User) (JWTOutput, error) {
	expirationTime := time.Now().Add(10 * time.Minute)
	claims := &Claims{
		UserId:   user.Id,
		Username: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
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
