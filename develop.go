package main

import (
	"fmt"
	"github.com/Azunyan1111/amazon/model"
	"time"
)

func main_() {
	hoge := make(chan int64)
	model.DataBaseInit()
	model.ApiInit()
	start := time.Now()
	// 1日一回実行するランキングWebスクレイピング関数。
	go func() { model.GetRankingASIN() }() //ok
	// 1回実行すればずっとASINから商品タイトルと画像URLを取得する関数
	go func() { model.GetItemInfoLoopForDatabases() }() //ok
	// 1日1回実行すると適当に拾ってきたASINリストから価格情報を取得して格納する。
	go func() { model.GetPrice() }()
	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
	defer model.MyDB.Close()
	<-hoge
}
