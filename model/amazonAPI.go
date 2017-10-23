package model

import (
	"fmt"
	"github.com/svvu/gomws/gmws"
	"github.com/svvu/gomws/mws/products"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/juju/errors"
	"strconv"
)



var client *products.Products

func ApiInit(){
	// API key config
	config := gmws.MwsConfig{
		SellerId:  os.Getenv("SellerId"),
		AccessKey: os.Getenv("AccessKey"),
		SecretKey: os.Getenv("SecretKey"),
		Region:    "JP",
	}
	var err error
	// Create client
	client, err = products.NewClient(config)
	if err != nil {
		log.Println(err)
		return
	}
}


// only go func
func GetItemInfoLoopForDatabases(){
	for _ = 1;;{

		hoge, err := SelectNotHaveInfoItemForASIN(5)
		if err != nil{
			fmt.Println(err)
		}
		if len(hoge) !=5{
			time.Sleep(1 * time.Hour)
		}
		start := time.Now()
		items, err := GetItemLookup(hoge)
		if err != nil{
			fmt.Println(err)
			time.Sleep(1 * time.Hour)
			continue
		}
		UpdateItemInfo(items)
		end := time.Now()
		// 1 H 18000 request
		if (end.Sub(start)).Seconds() < 1{
			time.Sleep(time.Second)
		}
	}
}
// only info
func GetItemLookup(asinList []string)([]Item, error){
	if len(asinList) != 5{
		return []Item{}, errors.New("asinList len err")
	}
	// send
	response1 := client.GetMatchingProductForId("ASIN",asinList)
	if response1.Error != nil || response1.StatusCode != http.StatusOK {
		return []Item{}, response1.Error
	}
	// responseXML to XMLNode
	xmlNode, _ := gmws.GenerateXMLNode(response1.Body)
	if gmws.HasErrors(xmlNode) {
		// TODO:正常にレスポンスは受け取っているが含まれていない謎のエラーが存在する
		//log.Println(gmws.GetErrors(xmlNode))
		//return []Item{}, errors.New("XML ERROR")
	}
	// set Item
	var items []Item
	for i, _ := range xmlNode.FindByKey("Title"){
		var item Item
		item.ASIN = xmlNode.FindByKey("ASIN")[i].Value.(string)
		item.Title = xmlNode.FindByKey("Title")[i].Value.(string)
		item.Image = xmlNode.FindByKey("URL")[i].Value.(string)
		items = append(items, item)
	}
	return items, nil
}

// go func only. 1 day
func GetPrice(){
	allASIN, err := SelectAllForASINLimit864000()
	if err != nil{
		log.Println(err)
		return
	}
	var asinDoubleArray [][]string
	var tempArray []string
	for i, asin := range allASIN {
		if i % 21 == 0{
			if len(tempArray) == 0{
				continue
			}
			asinDoubleArray = append(asinDoubleArray, tempArray)
			tempArray = []string{}
		}else {
			tempArray = append(tempArray,asin)
		}
	}

	// send api
	for _, asinArray := range asinDoubleArray{
		//*/
		start := time.Now()
		response := client.GetLowestOfferListingsForASIN(asinArray, gmws.Parameters{"ItemCondition":"New"})
		if response.Error != nil || response.StatusCode != http.StatusOK {
			log.Println("http Status:" + string(response.StatusCode))
			log.Println(response.Error)
			return
		}

		// responseXML to XMLNode
		xmlNode, _ := gmws.GenerateXMLNode(response.Body)
		if gmws.HasErrors(xmlNode) {
			//log.Println(gmws.GetErrors(xmlNode))
			//return
		}

		var saveProduct []ProductStock

		// Get all products
		products := xmlNode.FindByKey("GetLowestOfferListingsForASINResult")
		// products to one product
		for _, product := range products {
			// Get all stocks
			stocks := product.FindByPath("Product.LowestOfferListings.LowestOfferListing")
			// get data time
			insertTime := time.Now().Unix()

			// stocks to one stock
			for _, stock := range stocks {
				temp := ProductStock{
					ASIN:         product.FindByPath("Product.Identifiers.MarketplaceASIN.ASIN")[0].Value.(string),
					Amount:       stock.FindByPath("Price.LandedPrice.Amount")[0].Value.(string),
					Channel:      stock.FindByPath("Qualifiers.FulfillmentChannel")[0].Value.(string),
					Conditions:   stock.FindByPath("Qualifiers.ItemCondition")[0].Value.(string),
					ShippingTime: stock.FindByPath("Qualifiers.ShippingTime.Max")[0].Value.(string),
					InsertTime:   strconv.FormatInt(insertTime,10),
				}
				if !isInArray(saveProduct, temp){
					saveProduct = append(saveProduct, temp)
				}
			}
		}
		InsertProductPrice(saveProduct)

		end := time.Now()
		// 1 H 36000 request 1 request per 2second
		if (end.Sub(start)).Seconds() < 2{
			time.Sleep(2 * time.Second)
		}
		//*/
	}
}
func isInArray(s []ProductStock, e ProductStock) bool {
	for _, v := range s {
		if e.ASIN == v.ASIN {
			return true
		}
	}
	return false
}