package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"fmt"
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
func GetUrl() ([]string, error) {
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

func SetNewASIN(asins []string){
	for _, asin := range asins{
		_, err := MyDB.Exec("INSERT INTO Items(ASIN) VALUES(?)",asin)
		if err != nil {
			continue
		}
	}
}

func SetItemInfo(items []Item){
	for _, hoge := range items{
		_, err := MyDB.Exec("UPDATE Items SET title = ?, image = ? WHERE ASIN = ?",
			hoge.Title, hoge.Image, hoge.ASIN)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func GetItemNotHaveInfoASIN(limit int)([]string, error){
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
