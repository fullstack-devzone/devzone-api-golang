package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/domain"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string    `json:"access_token"`
	ExpiresAt   time.Time `json:"expires_at"`
	User        LoginUser `json:"user"`
}

type LoginUser struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func (a *AuthenticationController) Login(c *gin.Context) {
	ctx := c.Request.Context()
	var user LoginRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	loginUser, err := a.repository.Login(ctx, user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	jwtOutput, err := domain.CreateJwtToken(a.cfg, *loginUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, LoginResponse{
		AccessToken: jwtOutput.Token,
		ExpiresAt:   jwtOutput.Expires,
		User: LoginUser{
			Id:    loginUser.Id,
			Name:  loginUser.Name,
			Email: loginUser.Email,
			Role:  loginUser.Role,
		},
	})
}
