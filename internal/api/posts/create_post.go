package posts

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	log "github.com/sirupsen/logrus"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/domain"
)

type CreatePostModel struct {
	Title   string `json:"title" validate:"required"`
	Url     string `json:"url" validate:"required,url"`
	Content string `json:"content" validate:"required"`
}

func (post CreatePostModel) Validate() error {
	return validation.ValidateStruct(&post,
		validation.Field(&post.Title, validation.Required),
		validation.Field(&post.Url, validation.Required, is.URL),
		validation.Field(&post.Content, validation.Required),
	)
}

func (pc PostController) Create(c *gin.Context) {
	log.Info("create post")
	ctx := c.Request.Context()
	var createPost CreatePostModel
	if err := c.BindJSON(&createPost); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Unable to parse request body. Error: " + err.Error(),
		})
		return
	}
	err := createPost.Validate()
	if err != nil {
		log.Errorf("Error while create post %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to create post",
		})
		return
	}
	userId := c.MustGet("CurrentUserId").(int)
	now := time.Now()
	post := domain.Post{
		Title:       createPost.Title,
		Url:         createPost.Url,
		Content:     createPost.Content,
		CreatedBy:   userId,
		CreatedDate: &now,
	}
	post, err = pc.repository.CreatePost(ctx, post)
	if err != nil {
		log.Errorf("Error while create post %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to create post",
		})
		return
	}
	c.JSON(http.StatusCreated, post)
}
