package auth_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sivaprasadreddy/devzone-api-golang/cmd"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/api/auth"
	auth2 "github.com/sivaprasadreddy/devzone-api-golang/internal/auth"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/config"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/domain"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/test_support"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AuthControllerTestSuite struct {
	suite.Suite
	PgContainer *test_support.PostgresContainer
	cfg         config.AppConfig
	app         *cmd.App
	router      *gin.Engine
}

func (suite *AuthControllerTestSuite) SetupSuite() {
	suite.PgContainer = test_support.InitPostgresContainer()
	suite.cfg = config.GetConfig(".env")
	suite.app = cmd.NewApp(suite.cfg)
	suite.router = suite.app.Router
}

func (suite *AuthControllerTestSuite) TearDownSuite() {
	suite.PgContainer.CloseFn()
}

func TestAuthControllerTestSuite(t *testing.T) {
	suite.Run(t, new(AuthControllerTestSuite))
}

func (suite *AuthControllerTestSuite) TestValidLogin() {
	t := suite.T()
	reqBody := strings.NewReader(`
		{
			"username": "admin@gmail.com",
			"password": "admin"
		}
	`)
	req, _ := http.NewRequest(http.MethodPost, "/api/login", reqBody)
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	actualResponse := rr.Body
	var loginResponse auth.LoginResponse
	err := json.NewDecoder(actualResponse).Decode(&loginResponse)

	assert.Nil(t, err)
	assert.NotNil(t, loginResponse.AccessToken)
	assert.NotNil(t, loginResponse.ExpiresAt)
	assert.Equal(t, 1, loginResponse.User.Id)
	assert.Equal(t, "Admin", loginResponse.User.Name)
	assert.Equal(t, "admin@gmail.com", loginResponse.User.Email)
	assert.Equal(t, "ROLE_ADMIN", loginResponse.User.Role)
}

func (suite *AuthControllerTestSuite) TestInvalidLogin() {
	t := suite.T()
	reqBody := strings.NewReader(`
		{
			"username": "admin@gmail.com",
			"password": "invalidpassword"
		}
	`)
	req, _ := http.NewRequest(http.MethodPost, "/api/login", reqBody)
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
}

func (suite *AuthControllerTestSuite) TestGetCurrentUser() {
	t := suite.T()
	token, err := auth2.CreateJwtToken(suite.cfg, domain.User{
		Id:    1,
		Name:  "Admin",
		Email: "admin@gmail.com",
	})
	assert.Nil(t, err)

	req, _ := http.NewRequest(http.MethodGet, "/api/me", nil)
	req.Header.Add("Authorization", token.Token)
	rr := httptest.NewRecorder()
	suite.router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	actualResponse := rr.Body
	var userResponse auth.LoginUser
	err = json.NewDecoder(actualResponse).Decode(&userResponse)

	assert.Nil(t, err)
	assert.Equal(t, 1, userResponse.Id)
	assert.Equal(t, "Admin", userResponse.Name)
	assert.Equal(t, "admin@gmail.com", userResponse.Email)
}
