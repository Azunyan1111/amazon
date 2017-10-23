package model

type Item struct {
	ASIN string
	Title string
	Image string
}

type ProductStock struct {
	ASIN         string
	Amount       string
	Channel      string
	Conditions   string
	ShippingTime string
	InsertTime   int64
}