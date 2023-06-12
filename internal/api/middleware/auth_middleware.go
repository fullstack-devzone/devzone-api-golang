package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/config"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/domain"
)

func AuthMiddleware(cfg config.AppConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeaderValue := c.GetHeader("Authorization")
		if authHeaderValue == "" || !strings.HasPrefix(authHeaderValue, "Bearer ") {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenValue := authHeaderValue[len("Bearer "):]
		claims := &domain.Claims{}
		tkn, err := jwt.ParseWithClaims(tokenValue, claims,
			func(token *jwt.Token) (interface{}, error) {
				return []byte(cfg.JwtSecret), nil
			})
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if tkn == nil || !tkn.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set(config.AuthUserIdKey, claims.UserId)
		c.Next()
	}
}
