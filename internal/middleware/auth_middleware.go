package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/auth"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/config"
)

func AuthMiddleware(cfg config.AppConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenValue := c.GetHeader("Authorization")
		claims := &auth.Claims{}
		tkn, err := jwt.ParseWithClaims(tokenValue, claims,
			func(token *jwt.Token) (interface{}, error) {
				return []byte(cfg.JwtSecret), nil
			})
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		if tkn == nil || !tkn.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Set("CurrentUserId", claims.UserId)
		c.Next()
	}
}
