package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddDnsHandler(c echo.Context) error {

	return c.String(http.StatusOK, "Record added successfully")
}
