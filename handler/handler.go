package handler

import (
	"fmt"
	"net/http"

	"github.com/kouxi08/Artfolio/config"
	"github.com/kouxi08/Artfolio/pkg"
	"github.com/labstack/echo/v4"
)

func AddDnsHandler(c echo.Context) error {

	config := c.Get("config").(*config.Config)
	name := config.Name
	recordType := config.RecordType
	ttl := config.TTL
	content := config.Content

	//レコード追加
	resp, err := pkg.AddRecords(name, recordType, ttl, content)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	defer resp.Body.Close()
	fmt.Println("Response Status:", resp.Status)

	return c.String(http.StatusOK, "Record added successfully")
}
