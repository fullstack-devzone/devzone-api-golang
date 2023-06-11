package posts

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/domain"
)

type UpdatePostModel struct {
	Title   string `json:"title" binding:"required"`
	Url     string `json:"url" binding:"required,url"`
	Content string `json:"content" binding:"required"`
}

func (pc PostController) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	log.Infof("update post id=%d", id)
	ctx := c.Request.Context()
	var updatePost UpdatePostModel
	if err := c.ShouldBindJSON(&updatePost); err != nil {
		log.Errorf("Invalid request payload. Error: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload",
		})
		return
	}
	now := time.Now()
	var post = domain.Post{
		Id:          id,
		Title:       updatePost.Title,
		Url:         updatePost.Url,
		Content:     updatePost.Content,
		UpdatedDate: &now,
	}
	post, err := pc.repository.UpdatePost(ctx, post)
	if err != nil {
		log.Errorf("Failed to update post. Error: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update post",
		})
		return
	}
	postModel, _ := pc.repository.GetPostById(c.Request.Context(), id)
	c.JSON(http.StatusOK, postModel)
}
