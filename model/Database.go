package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"fmt"
	"github.com/juju/errors"
)

var MyDB *sql.DB

func DataBaseInit() {
	//hoge := "root:541279xx@tcp(mydbinstance.cv8ap3ddulzc.us-east-2.rds.amazonaws.com:3306)/amazon"
	dataSource := os.Getenv("DATABASE_URL")
	var err error
	MyDB, err = sql.Open("mysql", dataSource) //"root:@/my_database")
	if err != nil {
		panic(err)
	}
}

// get rank urls
// WANG this is 10 time second. only go func{}()
func SelectAllUrl() ([]string, error) {
	// TODO: LIMIT
	rows, err := MyDB.Query("SELECT URL FROM CategoryURL;")
	if err != nil {
		return nil, err
	}
	// list append
	var urls []string = make([]string, 0)
	for rows.Next() {
		var url string
		if err := rows.Scan(&url); err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}
	return urls, nil
}

func InsertNewASIN(asins []string){
	for _, asin := range asins{
		_, err := MyDB.Exec("INSERT INTO Items(ASIN) VALUES(?)",asin)
		if err != nil {
			continue
		}
	}
}

func UpdateItemInfo(items []Item){
	for _, hoge := range items{
		_, err := MyDB.Exec("UPDATE Items SET title = ?, image = ? WHERE ASIN = ?",
			hoge.Title, hoge.Image, hoge.ASIN)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func SelectNotHaveInfoItemForASIN(limit int)([]string, error){
	rows, err := MyDB.Query("SELECT ASIN FROM Items WHERE title IS null LIMIT ?", limit)
	if err != nil {
		return nil, err
	}
	// list append
	var asins []string = make([]string, 0)
	for rows.Next() {
		var asin string
		if err := rows.Scan(&asin); err != nil {
			return nil, err
		}
		asins = append(asins, asin)
	}
	return asins, nil
}

func SelectAllForASINLimit864000()([] string, error){
	// new connection
	dataSource := os.Getenv("DATABASE_URL")
	myDB, err := sql.Open("mysql", dataSource) //"root:@/my_database")
	if err != nil {
		return nil,errors.New("Can not connection Database")
	}

	// query. API MAX 86500 / day
	rows, err := MyDB.Query("SELECT ASIN FROM Items ORDER BY RAND() LIMIT 864000")
	if err != nil {
		return nil, err
	}
	// list append
	var asins []string
	for rows.Next() {
		var asin string
		if err := rows.Scan(&asin); err != nil {
			return nil, err
		}
		asins = append(asins, asin)
	}
	defer myDB.Close()
	return asins, nil
}

func InsertProductPrice(asins []ProductStock){
	for _, asin := range asins{
		_, err := MyDB.Exec("INSERT INTO Price(ASIN,Amount,Channel,Conditions,ShippingTime,InsertTime)" +
			" VALUES(?,?,?,?,?,?)",asin.ASIN,asin.Amount,asin.Channel,asin.Conditions,asin.ShippingTime,asin.InsertTime)
		if err != nil {
			continue
		}
	}
}

func SelectProductInfoForASIN(asin string)(Item, error){
	var product Item
	if err := MyDB.QueryRow("SELECT title,image FROM Items WHERE ASIN = ? LIMIT 1", asin).Scan(
		&product.Title,&product.Image); err != nil {
			return Item{ASIN:asin},err
	}
	product.ASIN = asin
	return product, nil
}

func SelectProductStockForASIN(asin string)([]ProductStock, error){
	// TODO: LIMIT
	rows, err := MyDB.Query("SELECT Amount,Channel,Conditions,ShippingTime," +
		"InsertTime FROM Price WHERE ASIN = ?;",asin)
	if err != nil {
		return []ProductStock{}, err
	}
	// list append
	var productStocks []ProductStock
	for rows.Next() {
		var productStock ProductStock
		if err := rows.Scan(&productStock.Amount,&productStock.Channel,
			&productStock.Conditions,&productStock.ShippingTime,&productStock.InsertTime); err != nil {
			return []ProductStock{}, err
		}
		productStock.ASIN = asin
		productStocks = append(productStocks, productStock)
	}
	return productStocks, nil
}