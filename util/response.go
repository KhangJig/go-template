package util

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

type response struct {
}

var Response response

func (response) Success(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "OK",
		"data":    data,
	})
}

func (response) Error(c echo.Context, httpCode int, err error) error {
	log.Error(err)

	return c.JSON(httpCode, map[string]interface{}{
		"code":    httpCode,
		"message": err.Error(),
	})
}
