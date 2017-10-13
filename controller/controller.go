package controller

import (
	"net/http"
	"github.com/labstack/echo"
)

// サイトで共通情報
type ServiceInfo struct {
	Title string
}

var serviceInfo = ServiceInfo {
	"Amazon",
}

type Product struct {
	Name string
}

type pageContentData struct{
	ServiceInfo
	Product
}

func MainPage() echo.HandlerFunc {
	//c をいじって Request, Responseを色々する
	return func(c echo.Context) error {
		// テンプレートに渡す値
		data := &pageContentData{serviceInfo,Product{"Name"}}
		return c.Render(http.StatusOK, "index", data)
	}
}