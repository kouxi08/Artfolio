package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func addDnsHandler(c echo.Context) error {
	// if err := Dns(c.Get("config").(*Config)); err != nil {
	// 	return err
	// }
	return c.String(http.StatusOK, "Record added successfully")
}
