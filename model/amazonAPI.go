package main

import (
	"fmt"
	"github.com/svvu/gomws/gmws"
	"github.com/svvu/gomws/mws/products"
	"log"
	"net/http"
	"os"
	"time"
)

type StockData struct {
	ASIN         string
	Amount       string
	Channel      string
	Condition    string
	ShippingTime string
	InsertTime   int64
}




func main() {
	// API key config
	config := gmws.MwsConfig{
		SellerId:  os.Getenv("SellerId"),
		AccessKey: os.Getenv("AccessKey"),
		SecretKey: os.Getenv("SecretKey"),
		Region:    "JP",
	}

	// Create client
	productsClient, err := products.NewClient(config)
	if err != nil {
		log.Println(err)
		return
	}

	// Send request
	response := productsClient.GetLowestOfferListingsForASIN([]string{"B00JPKHFTA", "B075VQ6RFZ"})
	if response.Error != nil || response.StatusCode != http.StatusOK {
		log.Println("http Status:" + string(response.StatusCode))
		log.Println(response.Error)
		return
	}

	// responseXML to XMLNode
	xmlNode, err := gmws.GenerateXMLNode(response.Body)
	if gmws.HasErrors(xmlNode) {
		log.Println(gmws.GetErrors(xmlNode))
		return
	}

	// Get all products
	products := xmlNode.FindByKey("GetLowestOfferListingsForASINResult")
	// products to one product
	for _, product := range products {
		// Get all stocks
		stocks := product.FindByPath("Product.LowestOfferListings.LowestOfferListing")

		insertTime := time.Now().Unix()

		// stocks to one stock
		for _, stock := range stocks {
			temp := StockData{
				ASIN:         product.FindByPath("Product.Identifiers.MarketplaceASIN.ASIN")[0].Value.(string),
				Amount:       stock.FindByPath("Price.LandedPrice.Amount")[0].Value.(string),
				Channel:      stock.FindByPath("Qualifiers.FulfillmentChannel")[0].Value.(string),
				Condition:    stock.FindByPath("Qualifiers.ItemCondition")[0].Value.(string),
				ShippingTime: stock.FindByPath("Qualifiers.ShippingTime.Max")[0].Value.(string),
				InsertTime:   insertTime,
			}

			fmt.Println(temp)
		}
	}
}
