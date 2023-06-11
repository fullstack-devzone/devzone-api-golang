package posts

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	log "github.com/sirupsen/logrus"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/domain"
)

type UpdatePostModel struct {
	Title   string `json:"title"`
	Url     string `json:"url"`
	Content string `json:"content"`
}

func (l UpdatePostModel) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Title, validation.Required),
		validation.Field(&l.Url, validation.Required, is.URL),
		validation.Field(&l.Content, validation.Required),
	)
}

func (b PostController) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	log.Infof("update post id=%d", id)
	ctx := c.Request.Context()
	var updatePost UpdatePostModel
	if err := c.BindJSON(&updatePost); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Unable to parse request body. Error: " + err.Error(),
		})
		return
	}
	err := updatePost.Validate()
	if err != nil {
		log.Errorf("Error while update post %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to update post",
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
	post, err = b.repository.UpdatePost(ctx, post)
	if err != nil {
		log.Errorf("Error while update post")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Unable to update post",
		})
		return
	}
	postModel, _ := b.repository.GetPostById(c.Request.Context(), id)
	c.JSON(http.StatusOK, postModel)
}
