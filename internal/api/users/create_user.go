package users

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	log "github.com/sirupsen/logrus"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/domain"
)

type CreateUserModel struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (l CreateUserModel) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Name, validation.Required),
		validation.Field(&l.Email, validation.Required, is.Email),
		validation.Field(&l.Password, validation.Required),
	)
}

func (b UserController) Create(c *gin.Context) {
	log.Info("create user")
	ctx := c.Request.Context()
	var createUser CreateUserModel
	if err := c.BindJSON(&createUser); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Unable to parse request body. Error: " + err.Error(),
		})
		return
	}
	err := createUser.Validate()
	if err != nil {
		log.Errorf("Error while create user %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to create user",
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
	user, err = b.repository.CreateUser(ctx, user)
	if err != nil {
		log.Errorf("Error while create user %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to create user",
		})
		return
	}
	c.JSON(http.StatusCreated, user)
}
