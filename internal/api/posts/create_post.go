package posts

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/config"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/domain"
)

type CreatePostModel struct {
	Title   string `json:"title" binding:"required"`
	Url     string `json:"url" binding:"required,url"`
	Content string `json:"content" binding:"required"`
}

func (pc PostController) Create(c *gin.Context) {
	log.Info("create post")
	ctx := c.Request.Context()
	var createPost CreatePostModel
	if err := c.ShouldBindJSON(&createPost); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Unable to parse request body. Error: " + err.Error(),
		})
		return
	}
	userId := c.MustGet(config.AuthUserIdKey).(int)
	now := time.Now()
	post := domain.Post{
		Title:       createPost.Title,
		Url:         createPost.Url,
		Content:     createPost.Content,
		CreatedBy:   userId,
		CreatedDate: &now,
	}
	post, err := pc.repository.CreatePost(ctx, post)
	if err != nil {
		log.Errorf("Error while create post %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to create post",
		})
		return
	}
	c.JSON(http.StatusCreated, post)
}
