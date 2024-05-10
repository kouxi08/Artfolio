package main

import (
	"github.com/kouxi08/Artfolio/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	//jsonファイルのデコード
	//サーバ起動
	server()
}

func server() {
	//インスタンス作成
	e := echo.New()

	//ミドルウェア設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	//レコード追加処理へ
	e.POST("/", handler.CreateHandler)
	//レコード削除処理へ
	e.PATCH("/", handler.DeleteHandler)
	//minioバケット追加処理
	e.POST("/minio", handler.MakeBucketHandler)

	e.Logger.Fatal(e.Start(":8088"))
}
