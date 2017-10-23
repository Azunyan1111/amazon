package main

import (
	"github.com/Azunyan1111/amazon/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"os"
	"github.com/Azunyan1111/amazon/model"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	// 初期処理
	model.DataBaseInit()
	model.ApiInit()

	// Echoのインスタンス作る
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// HTML読み込み
	t := &Template{templates: template.Must(template.ParseGlob("views/*.html"))}

	e.Renderer = t

	// ルーティング
	e.GET("/", controller.MainPage())
	e.GET("/:asin", controller.ProductPage())

	e.Static("/assets", "assets")

	defer model.MyDB.Close()
	// サーバー起動
	e.Start(":" + os.Getenv("PORT"))
}
