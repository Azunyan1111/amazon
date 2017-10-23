package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/Azunyan1111/amazon/model"
	"fmt"
	"strings"
)


// サイト共有情報記入
var serviceInfo = model.ServiceInfo{
	"Amazon",
}

func MainPage() echo.HandlerFunc {
	//c をいじって Request, Responseを色々する
	return func(c echo.Context) error {
		// テンプレートに渡す値
		data := &model.PageContentData{ServiceInfo:serviceInfo, Item:model.Item{}}
		return c.Render(http.StatusOK, "index", data)
	}
}

func ProductPage() echo.HandlerFunc {
	//c をいじって Request, Responseを色々する
	return func(c echo.Context) error {
		// 商品ASIN取得
		item, err := model.SelectProductInfoForASIN(c.Param("asin"))
		// 商品情報取得
		if err != nil{
			//TODO: 商品が登録されていいないorそもそも存在しない場合もあるため500は不適切である可能性がある。
			fmt.Println(err)
			data := &model.PageContentData{ServiceInfo:serviceInfo,Message:"500 InternalServerError. " +
				"サーバ内部エラーです。", Item:model.Item{}}
			return c.Render(http.StatusInternalServerError, "error", data)
		}
		// 商品在庫取得
		productStocks, err := model.SelectProductStockForASIN(item.ASIN)
		if err != nil{
			// データが存在しない場合も含む
			fmt.Println(err)
			//data := &model.PageContentData{ServiceInfo:serviceInfo,Message:"500 InternalServerError. " +
			//	"サーバ内部エラーです。", Item:model.Item{}}
			//return c.Render(http.StatusInternalServerError, "error", data)
		}

		// タストルカスタマイズ
		customServiceInfo := serviceInfo
		customServiceInfo.Title = serviceInfo.Title + " | " + item.Title
		// 画像画質変更
		item.Image = strings.Replace(item.Image, "SL75", "SL1500", 1)
		data := &model.PageContentData{ServiceInfo:customServiceInfo, Item: item, ProductStocks: productStocks}
		return c.Render(http.StatusOK, "index2", data)
	}
}