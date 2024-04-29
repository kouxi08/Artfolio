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
	//jsonファイルのデコード
	config, err := LoadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}
	//サーバ起動
	server(config)
}

func env() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
}

func server(config *Config) {
	//インスタンス作成
	e := echo.New()

	//ミドルウェア設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("config", config)
			return next(c)
		}
	})
	//レコード追加処理にルーティング
	e.GET("/addrecode", addDnsHandler)
	// e.GET("/", showDnsHandler)

	e.Logger.Fatal(e.Start(":8088"))
}

func addDnsHandler(c echo.Context) error {
	if err := dns(c.Get("config").(*Config)); err != nil {
		return err
	}
	return c.String(http.StatusOK, "Record added successfully")
}

func dns(config *Config) error {
	env()

	api := API{
		Endpoint: os.Getenv("ENDPOINT"),
		ApiKey:   os.Getenv("APIKEY"),
	}

	name := config.Name
	recordType := config.RecordType
	ttl := config.TTL
	content := config.Content

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
