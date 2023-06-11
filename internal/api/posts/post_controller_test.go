package posts_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sivaprasadreddy/devzone-api-golang/cmd"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/config"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/domain"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/test_support"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PostControllerTestSuite struct {
	suite.Suite
	PgContainer *test_support.PostgresContainer
	cfg         config.AppConfig
	app         *cmd.App
	router      *gin.Engine
}

func (suite *PostControllerTestSuite) SetupSuite() {
	suite.PgContainer = test_support.InitPostgresContainer()
	suite.cfg = config.GetConfig(".env")
	suite.app = cmd.NewApp(suite.cfg)
	suite.router = suite.app.Router
}

func (suite *PostControllerTestSuite) TearDownSuite() {
	suite.PgContainer.CloseFn()
}

func TestPostControllerTestSuite(t *testing.T) {
	suite.Run(t, new(PostControllerTestSuite))
}

func (suite *PostControllerTestSuite) TestGetAllPosts() {
	t := suite.T()
	req, _ := http.NewRequest(http.MethodGet, "/api/posts", nil)
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	actualResponse := rr.Body.String()
	assert.NotEqual(t, "[]", actualResponse,
		"Expected an non-empty array. Got %s", actualResponse)
}

func (suite *PostControllerTestSuite) TestGetPostById() {
	t := suite.T()
	req, _ := http.NewRequest(http.MethodGet, "/api/posts/1", nil)
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	actualResponse := rr.Body
	var postResponse domain.PostModel
	err := json.NewDecoder(actualResponse).Decode(&postResponse)

	assert.Nil(t, err)
	assert.NotNil(t, postResponse.Id)
	assert.Equal(t, "How To Remove Docker Containers, Images, Volumes, and Networks", postResponse.Title)
	assert.Equal(t, "https://linuxize.com/post/how-to-remove-docker-images-containers-volumes-and-networks/", postResponse.Url)
	assert.Equal(t, "How To Remove Docker Containers, Images, Volumes, and Networks", postResponse.Content)
	assert.Equal(t, 1, postResponse.CreatedBy.Id)
	assert.Equal(t, "Admin", postResponse.CreatedBy.Name)
	assert.Equal(t, "admin@gmail.com", postResponse.CreatedBy.Email)
}

func (suite *PostControllerTestSuite) TestCreatePost() {
	t := suite.T()
	reqBody := strings.NewReader(`
		{
			"title": "Test Post title",
			"url":     "https://example.com",
			"content": "Test Post content"
		}
	`)
	token, err := domain.CreateJwtToken(suite.cfg, domain.User{
		Id:    1,
		Name:  "Siva",
		Email: "siva@gmail.com",
	})
	assert.Nil(t, err)

	req, _ := http.NewRequest(http.MethodPost, "/api/posts", reqBody)
	req.Header.Add("Authorization", token.Token)
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	actualResponse := rr.Body
	var postResponse domain.Post
	err = json.NewDecoder(actualResponse).Decode(&postResponse)

	assert.Nil(t, err)
	assert.NotNil(t, postResponse.Id)
	assert.Equal(t, "Test Post title", postResponse.Title)
	assert.Equal(t, "https://example.com", postResponse.Url)
	assert.Equal(t, "Test Post content", postResponse.Content)
}

func (suite *PostControllerTestSuite) TestUpdatePost() {
	t := suite.T()
	reqBody := strings.NewReader(`
		{
			"title": "Test Post title",
			"url":     "https://example.com",
			"content": "Test Post content"
		}
	`)
	token, err := domain.CreateJwtToken(suite.cfg, domain.User{
		Id:    1,
		Name:  "Siva",
		Email: "siva@gmail.com",
	})
	assert.Nil(t, err)

	req, _ := http.NewRequest(http.MethodPut, "/api/posts/1", reqBody)
	req.Header.Add("Authorization", token.Token)
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	actualResponse := rr.Body
	var postResponse domain.PostModel
	err = json.NewDecoder(actualResponse).Decode(&postResponse)

	assert.Nil(t, err)
	assert.Equal(t, 1, postResponse.Id)
	assert.Equal(t, "Test Post title", postResponse.Title)
	assert.Equal(t, "https://example.com", postResponse.Url)
	assert.Equal(t, "Test Post content", postResponse.Content)
}

func (suite *PostControllerTestSuite) TestDeletePost() {
	t := suite.T()
	token, err := domain.CreateJwtToken(suite.cfg, domain.User{
		Id:    1,
		Name:  "Siva",
		Email: "siva@gmail.com",
	})
	assert.Nil(t, err)

	req, _ := http.NewRequest(http.MethodDelete, "/api/posts/2", nil)
	req.Header.Add("Authorization", token.Token)
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
