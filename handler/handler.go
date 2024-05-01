package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kouxi08/Artfolio/config"
	"github.com/kouxi08/Artfolio/pkg"
	"github.com/labstack/echo/v4"
)

func CreateHandler(c echo.Context) error {
	config, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}

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

	//deployment作成
	pkg.CreateDeployment(siteName, deploymentName)
	//service作成
	pkg.CreateService(siteName, serviceName)
	//ingress作成
	pkg.CreateIngress(ingressName, hostName, serviceName)

	return c.String(http.StatusOK, "Record added successfully")
}

func DeleteHandler(c echo.Context) error {
	config, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}

	siteName := c.FormValue("name")
	deleteRecordName := fmt.Sprintf("%s%s", siteName, config.Record.Name)

	deploymentName := fmt.Sprintf("%s%s", siteName, config.KubeConfig.DeploymentName)
	serviceName := fmt.Sprintf("%s%s", siteName, config.KubeConfig.ServiceName)
	ingressName := fmt.Sprintf("%s%s", siteName, config.KubeConfig.IngressName)

	resp, err := pkg.DeleteRecords(deleteRecordName)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	defer resp.Body.Close()
	fmt.Println("Response Status:", resp.Status)

	pkg.DeleteDeployment(deploymentName)
	pkg.DeleteService(serviceName)
	pkg.DeleteIngress(ingressName)
	return c.String(http.StatusOK, "Record delete successfully")
}
