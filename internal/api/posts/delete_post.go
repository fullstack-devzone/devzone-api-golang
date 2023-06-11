package posts

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (pc PostController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	log.Infof("delete post with id=%d", id)
	ctx := c.Request.Context()
	err := pc.repository.DeletePost(ctx, id)
	if err != nil {
		log.Errorf("Error while deleting post")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to delete post",
		})
		return
	}
	c.JSON(http.StatusOK, nil)
}
