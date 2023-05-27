package posts_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllPosts(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/posts", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	actualResponseJson := rr.Body.String()
	assert.NotEqual(t, "[]", actualResponseJson)
	assert.NotEqual(t, actualResponseJson, "[]",
		"Expected an non-empty array. Got %s", actualResponseJson)
}
