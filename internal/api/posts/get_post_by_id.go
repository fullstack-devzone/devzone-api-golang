package posts

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/domain"
)

func (pc PostController) GetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	log.Infof("Fetching post by id %d", id)
	ctx := c.Request.Context()
	post, err := pc.repository.GetPostById(ctx, id)
	if err != nil {
		log.Errorf("Error while fetching post by id: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to fetch post by id",
		})
		return
	}
	c.JSON(http.StatusOK, domain.PostModel{
		Id:      post.Id,
		Title:   post.Title,
		Url:     post.Url,
		Content: post.Content,
		CreatedBy: domain.PostCreatedByModel{
			Id:    post.CreatedBy.Id,
			Name:  post.CreatedBy.Name,
			Email: post.CreatedBy.Email,
		},
		CreatedDate: post.CreatedDate,
		UpdatedDate: post.UpdatedDate,
	})
}
