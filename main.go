package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type API struct {
	Endpoint string
	ApiKey   string
}

func main() {
	//インスタンス作成
	e := echo.New()

	//ミドルウェア設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", dns)

	e.Logger.Fatal(e.Start(":8088"))
}

func dns(c echo.Context) error {
	if err := sub(); err != nil {
		return err
	}
	return c.String(http.StatusOK, "Record added successfully")
}

func sub() error {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	api := API{
		Endpoint: os.Getenv("ENDPOINT"),
		ApiKey:   os.Getenv("APIKEY"),
	}

	name := "example.hoge.com."
	recordType := "CNAME"
	ttl := "3600"
	content := "hoge.hoge.com."

	//ゾーンの参照
	// showZones(api)
	//レコード追加
	resp, err := addRecords(api, name, recordType, ttl, content)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	defer resp.Body.Close()
	fmt.Println("Response Status:", resp.Status)

	return nil
}
