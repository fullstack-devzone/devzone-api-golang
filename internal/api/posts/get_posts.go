package posts

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/domain"
)

func (pc PostController) GetPosts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	query := c.DefaultQuery("query", "")
	log.Infof("Fetching posts with query: %s and page: %d", query, page)
	ctx := c.Request.Context()
	var postsPage domain.PostsPageModel
	var err error
	if query == "" {
		postsPage, err = pc.repository.GetPosts(ctx, page)
	} else {
		postsPage, err = pc.repository.SearchPosts(ctx, query, page)
	}
	if err != nil {
		log.Errorf("Error while fetching posts: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to fetch posts",
		})
		return
	}
	c.JSON(http.StatusOK, postsPage)
}
