package handler

import (
	"fmt"
	"myapp/08_rest_api_example/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	userID := c.Param("id")
	id := 0
	if _, err := fmt.Sscanf(userID, "%d", &id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	user, ok := model.Users[id]
	if !ok {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	return c.JSON(http.StatusOK, user)
}

func GetAllUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, model.Users)
}
