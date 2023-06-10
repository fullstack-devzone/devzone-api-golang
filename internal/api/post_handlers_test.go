package api_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sivaprasadreddy/devzone-api-golang/cmd"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/config"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/test_support"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PostApisTestSuite struct {
	suite.Suite
	PgContainer *test_support.PostgresContainer
	cfg         config.AppConfig
	app         *cmd.App
	router      *gin.Engine
}

func (suite *PostApisTestSuite) SetupSuite() {
	fmt.Println("-----------SetupSuite()-----------")
	suite.PgContainer = test_support.InitPostgresContainer()
	suite.cfg = config.GetConfig(".env")
	suite.app = cmd.NewApp(suite.cfg)
	suite.router = suite.app.Router
}

func (suite *PostApisTestSuite) TearDownSuite() {
	fmt.Println("-----------TearDownSuite()-----------")
	suite.PgContainer.CloseFn()
}

func TestPostApisTestSuite(t *testing.T) {
	suite.Run(t, new(PostApisTestSuite))
}

func (suite *PostApisTestSuite) TestGetAllPosts() {
	t := suite.T()
	req, _ := http.NewRequest("GET", "/api/posts", nil)
	req.Header.Add("X-API-KEY", "1234")
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	actualResponseJson := rr.Body.String()
	assert.NotEqual(t, "[]", actualResponseJson)
	assert.NotEqual(t, actualResponseJson, "[]",
		"Expected an non-empty array. Got %s", actualResponseJson)
}
