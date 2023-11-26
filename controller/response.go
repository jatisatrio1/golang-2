package controller

import (
	"net/http"
	"rest-api/model"

	"github.com/labstack/echo"
)

func SetSuccessResponse(c echo.Context, message string, data interface{}) error {
	return c.JSON(http.StatusOK, model.Response{
		Message: message,
		Success: true,
		Data:    data,
	})
}

func SetErrorResponse(c echo.Context, message string) error {
	return c.JSON(http.StatusOK, model.Response{
		Message: message,
		Success: false,
	})
}
