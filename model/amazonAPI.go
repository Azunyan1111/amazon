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
)

type StockData struct {
	ASIN         string
	Amount       string
	Channel      string
	Condition    string
	ShippingTime string
	InsertTime   int64
}

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

		hoge, err := GetItemNotHaveInfoASIN(5)
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
		SetItemInfo(items)
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