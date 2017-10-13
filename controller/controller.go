package controller

import (
	"net/http"
	"github.com/labstack/echo"
)

// サイトで共通情報
type ServiceInfo struct {
	Title string
}
// 商品の情報
type Product struct {
	Name string
}
// コンテンツHTMLに渡す情報
type pageContentData struct{
	ServiceInfo
	Product
}
// サイト共有情報記入
var serviceInfo = ServiceInfo {
	"Amazon",
}


func MainPage() echo.HandlerFunc {
	//c をいじって Request, Responseを色々する
	return func(c echo.Context) error {
		// テンプレートに渡す値
		data := &pageContentData{serviceInfo,Product{"Name"}}
		return c.Render(http.StatusOK, "index", data)
	}
}