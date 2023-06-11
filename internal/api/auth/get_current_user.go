package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a AuthenticationController) GetCurrentUser(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.MustGet("CurrentUserId").(int)

	loginUser, err := a.repository.GetUserById(ctx, userId)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorised"})
		return
	}
	c.JSON(http.StatusOK, LoginUser{
		Id:    loginUser.Id,
		Name:  loginUser.Name,
		Email: loginUser.Email,
		Role:  loginUser.Role,
	})
}
