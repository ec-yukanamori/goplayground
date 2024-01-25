package handler_test

import (
	"myapp/08_rest_api_example/handler"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Test with an existing user
	if assert.NoError(t, handler.GetUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "\"id\":1")
		assert.Contains(t, rec.Body.String(), "\"name\":\"John Doe\"")
	}

	// Test with a non-existing user
	req = httptest.NewRequest(http.MethodGet, "/", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("99")

	if assert.NoError(t, handler.GetUser(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.Contains(t, rec.Body.String(), "\"error\":\"User not found\"")
	}

	// Test with an invalid user ID
	req = httptest.NewRequest(http.MethodGet, "/", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("invalid")

	if assert.NoError(t, handler.GetUser(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Contains(t, rec.Body.String(), "\"error\":\"Invalid user ID\"")
	}
}

func TestGetAllUsers(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test
	if assert.NoError(t, handler.GetAllUsers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "\"id\":1")
		assert.Contains(t, rec.Body.String(), "\"name\":\"John Doe\"")
		assert.Contains(t, rec.Body.String(), "\"id\":2")
		assert.Contains(t, rec.Body.String(), "\"name\":\"Jane Doe\"")
	}
}
