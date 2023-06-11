package posts

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/domain"
)

func (b PostController) GetAll(c *gin.Context) {
	log.Info("Fetching all posts")
	ctx := c.Request.Context()
	posts, err := b.repository.GetPosts(ctx)
	if err != nil {
		log.Errorf("Error while fetching posts")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to fetch posts",
		})
		return
	}
	if posts == nil {
		posts = []domain.Post{}
	}
	c.JSON(http.StatusOK, posts)
}
