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
		log.Errorf("Failed to delete post. Error: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete post",
		})
		return
	}
	c.JSON(http.StatusOK, nil)
}
