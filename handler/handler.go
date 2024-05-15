package handler

import (
	"fmt"
	"net/http"

	"github.com/kouxi08/Artfolio/config"
	"github.com/kouxi08/Artfolio/pkg"
	"github.com/labstack/echo/v4"
)

func CreateHandler(c echo.Context) error {
	config, _ := config.LoadConfig("config.json")

	siteName := c.FormValue("name")

	newRecordName := fmt.Sprintf("%s%s", siteName, config.Record.Name)
	deploymentName := fmt.Sprintf("%s%s", siteName, config.KubeConfig.DeploymentName)
	serviceName := fmt.Sprintf("%s%s", siteName, config.KubeConfig.ServiceName)
	ingressName := fmt.Sprintf("%s%s", siteName, config.KubeConfig.IngressName)
	hostName := fmt.Sprintf("%s%s", siteName, config.KubeConfig.HostName)

	pkg.AddDnsResouces(newRecordName, config.Record.RecordType, config.Record.TTL, config.Record.Content)
	pkg.CreateResources(siteName, deploymentName, serviceName, ingressName, hostName)

	return c.String(http.StatusOK, "Record added successfully")
}

func DeleteHandler(c echo.Context) error {
	config, _ := config.LoadConfig("config.json")

	siteName := c.FormValue("name")

	deleteRecordName := fmt.Sprintf("%s%s", siteName, config.Record.Name)
	deploymentName := fmt.Sprintf("%s%s", siteName, config.KubeConfig.DeploymentName)
	serviceName := fmt.Sprintf("%s%s", siteName, config.KubeConfig.ServiceName)
	ingressName := fmt.Sprintf("%s%s", siteName, config.KubeConfig.IngressName)

	pkg.DeleteDnsResouces(deleteRecordName)
	pkg.DeleteResources(deploymentName, serviceName, ingressName)

	return c.String(http.StatusOK, "Record delete successfully")
}

func MakeBucketHandler(c echo.Context) error {
	bucketName := c.FormValue("name")
	message, err := pkg.MakeBucket(bucketName)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error()) // エラーをHTTPレスポンスとして返す
	}
	fmt.Print(message)
	return c.JSON(http.StatusOK, map[string]interface{}{"bkname": message})
}
