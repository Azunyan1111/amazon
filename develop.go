package main

import (
	"github.com/Azunyan1111/amazon/model"
	"time"
	"fmt"
)

func main() {
	hoge := make(chan int64)
	model.DataBaseInit()
	model.ApiInit()
	start := time.Now()
	// 1日一回実行するランキングWebスクレイピング関数。
	go func() {model.GetRankingASIN()}() //ok
	// 1回実行すればずっとASINから商品タイトルと画像URLを取得する関数
	go func() {model.GetItemInfoLoopForDatabases()}() //ok
	end := time.Now()
	fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())
	defer model.MyDB.Close()
	<-hoge
}
