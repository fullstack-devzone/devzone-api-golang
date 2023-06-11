package users

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/domain"
)

type CreateUserModel struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (uc UserController) Create(c *gin.Context) {
	log.Info("create user")
	ctx := c.Request.Context()
	var createUser CreateUserModel
	if err := c.ShouldBindJSON(&createUser); err != nil {
		log.Errorf("Error in parsing create user payload: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload",
		})
		return
	}
	now := time.Now()
	user := domain.User{
		Name:        createUser.Name,
		Email:       createUser.Email,
		Password:    createUser.Password,
		Role:        "ROLE_USER",
		CreatedDate: &now,
	}
	user, err := uc.repository.CreateUser(ctx, user)
	if err != nil {
		log.Errorf("Error while creating user: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}
	c.JSON(http.StatusCreated, user)
}
