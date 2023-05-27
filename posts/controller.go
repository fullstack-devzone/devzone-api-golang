package posts

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type PostController struct {
	repository PostRepository
}

func NewPostController(repository PostRepository) *PostController {
	return &PostController{repository}
}

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
		posts = []Post{}
	}
	c.JSON(http.StatusOK, posts)
}

func (b PostController) GetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	log.Infof("Fetching post by id %d", id)
	ctx := c.Request.Context()
	post, err := b.repository.GetPostById(ctx, id)
	if err != nil {
		log.Errorf("Error while fetching post by id")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to fetch post by id",
		})
		return
	}
	c.JSON(http.StatusOK, post)
}

func (b PostController) Create(c *gin.Context) {
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
	post := Post{
		Title:       createPost.Title,
		Url:         createPost.Url,
		CreatedDate: time.Time{},
	}
	post, err = b.repository.CreatePost(ctx, post)
	if err != nil {
		log.Errorf("Error while create post %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to create post",
		})
		return
	}
	c.JSON(http.StatusCreated, post)
}

func (b PostController) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	log.Infof("update post id=%d", id)
	ctx := c.Request.Context()
	var post Post
	if err := c.BindJSON(&post); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Unable to parse request body. Error: " + err.Error(),
		})
		return
	}
	post.Id = id
	post.UpdatedDate = time.Now()
	post, err := b.repository.UpdatePost(ctx, post)
	if err != nil {
		log.Errorf("Error while update post")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Unable to update post",
		})
		return
	}
	post, _ = b.repository.GetPostById(c.Request.Context(), id)
	c.JSON(http.StatusOK, post)
}

func (b PostController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	log.Infof("delete post with id=%d", id)
	ctx := c.Request.Context()
	err := b.repository.DeletePost(ctx, id)
	if err != nil {
		log.Errorf("Error while deleting post")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to delete post",
		})
		return
	}
	c.JSON(http.StatusOK, nil)
}
