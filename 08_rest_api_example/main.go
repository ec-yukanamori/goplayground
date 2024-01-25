package main

import (
	"myapp/08_rest_api_example/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	// Echoインスタンスの作成
	e := echo.New()

	// ルーティング
	e.GET("/users/:id", handler.GetUser)
	e.GET("/users", handler.GetAllUsers)

	// サーバーの起動
	e.Start(":1323")
}
