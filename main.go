package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/Azunyan1111/amazon/controller"
	"io"
	"html/template"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	// Echoのインスタンス作る
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// HTML読み込み
	t := &Template{templates: template.Must(template.ParseGlob("views/*.html")),}
	e.Renderer = t

	// ルーティング
	e.GET("/", controller.MainPage())
	e.Static("/assets", "assets")

	// サーバー起動
	e.Start(":8888")
}
