package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (b UserController) GetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	log.Infof("Fetching user by id %d", id)
	ctx := c.Request.Context()
	post, err := b.repository.GetUserById(ctx, id)
	if err != nil {
		log.Errorf("Error while fetching user by id: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to fetch user by id",
		})
		return
	}
	c.JSON(http.StatusOK, post)
}
