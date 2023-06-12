package users_test

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

type UserControllerTestSuite struct {
	suite.Suite
	PgContainer *test_support.PostgresContainer
	cfg         config.AppConfig
	app         *cmd.App
	router      *gin.Engine
}

func (suite *UserControllerTestSuite) SetupSuite() {
	suite.PgContainer = test_support.InitPostgresContainer()
	suite.cfg = config.GetConfig(".env")
	suite.app = cmd.NewApp(suite.cfg)
	suite.router = suite.app.Router
}

func (suite *UserControllerTestSuite) TearDownSuite() {
	suite.PgContainer.CloseFn()
}

func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}

func (suite *UserControllerTestSuite) TestCreateUser() {
	t := suite.T()
	reqBody := strings.NewReader(`
		{
			"name":"newuser",
			"email": "newuser@gmail.com",
			"password": "secret"
		}
	`)

	req, _ := http.NewRequest(http.MethodPost, "/api/users", reqBody)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var userResponse domain.User
	err := json.NewDecoder(w.Body).Decode(&userResponse)

	assert.Nil(t, err)
	assert.NotNil(t, userResponse.Id)
	assert.Equal(t, "newuser", userResponse.Name)
	assert.Equal(t, "newuser@gmail.com", userResponse.Email)
	assert.Equal(t, "ROLE_USER", userResponse.Role)
}

func (suite *UserControllerTestSuite) TestGetUserById() {
	t := suite.T()
	req, _ := http.NewRequest(http.MethodGet, "/api/users/1", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var userResponse domain.User
	err := json.NewDecoder(w.Body).Decode(&userResponse)

	assert.Nil(t, err)
	assert.Equal(t, 1, userResponse.Id)
	assert.Equal(t, "Admin", userResponse.Name)
	assert.Equal(t, "admin@gmail.com", userResponse.Email)
	assert.Equal(t, "ROLE_ADMIN", userResponse.Role)
}
