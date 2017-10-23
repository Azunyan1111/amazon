package model

// サイトで共通情報
type ServiceInfo struct {
	Title string
}

// コンテンツHTMLに渡す情報
type PageContentData struct {
	ServiceInfo
	Message    string
	SubMessage string
	Item
	ProductStocks []ProductStock
}

type Item struct {
	ASIN  string
	Title string
	Image string
}

type ProductStock struct {
	ASIN         string
	Amount       string
	Channel      string
	Conditions   string
	ShippingTime string
	InsertTime   string
}
