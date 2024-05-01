package handler

import (
	"fmt"
	"net/http"

	"github.com/kouxi08/Artfolio/config"
	"github.com/kouxi08/Artfolio/pkg"
	"github.com/labstack/echo/v4"
)

func CreateHandler(c echo.Context) error {
	config := c.Get("config").(*config.Config)

	siteName := c.FormValue("name")
	newRecordName := fmt.Sprintf("%s%s", siteName, config.Record.Name)

	deploymentName := fmt.Sprintf("%s%s", siteName, config.KubeConfig.DeploymentName)
	serviceName := fmt.Sprintf("%s%s", siteName, config.KubeConfig.ServiceName)
	ingressName := fmt.Sprintf("%s%s", siteName, config.KubeConfig.IngressName)
	hostName := fmt.Sprintf("%s%s", siteName, config.KubeConfig.HostName)

	//レコード追加
	resp, err := pkg.AddRecords(newRecordName, config.Record.RecordType, config.Record.TTL, config.Record.Content)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	defer resp.Body.Close()
	fmt.Println("Response Status:", resp.Status)

	pkg.CreateDeployment(siteName, deploymentName)
	pkg.CreateService(siteName, serviceName)
	pkg.CreateIngress(ingressName, hostName, serviceName)

	return c.String(http.StatusOK, "Record added successfully")
}

func DeleteHandler(c echo.Context) error {
	config := c.Get("config").(*config.Config)

	siteName := c.FormValue("name")

	name := fmt.Sprintf("%s%s", siteName, config.Record.Name)

	resp, err := pkg.DeleteRecords(name)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	defer resp.Body.Close()
	fmt.Println("Response Status:", resp.Status)
	return c.String(http.StatusOK, "Record delete successfully")
}
