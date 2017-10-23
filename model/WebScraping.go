package model

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/Songmu/retry"
	"github.com/pkg/errors"
	"log"
	"regexp"
	"strconv"
	"time"
)

var r = regexp.MustCompile(`[A-Z0-9]{10}/`)

// チャンルに値が入ったら処理が実行される
func sendWebScraping(myCh chan string) {
	// myChからの送信を受け取る
	for out := range myCh {
		// 処理
		asins, err := getRank(out)
		if err != nil {
			fmt.Println(err.Error())
		}
		InsertNewASIN(asins)
	}
}

// only go func{}()
func GetRankingASIN() {

	// intを送受信どちらもできるchannel
	myChan := make(chan string)

	endChan := make(chan int)

	// 並行処理goroutineを立ち上げる
	go sendWebScraping(myChan)

	//myChan <- "hoge" // 100を流し込む

	// get urls
	urls, err := SelectAllUrl()
	if err != nil {
		log.Println(err)
		return
	}

	connectPerSecond := (len(urls) / 1440) + 2
	go func() {
		for i := 0; i < len(urls); i++ {
			if i%connectPerSecond == 0 {
				time.Sleep(60 * time.Second)
			}
			myChan <- urls[i]
		}
		endChan <- 114514
	}()
	// 終了するまでまつ。
	<-endChan
}

func getRank(url string) ([]string, error) {
	// 返すASIN
	var asin []string
	var doc *goquery.Document
	var err error

	err = retry.Retry(100, 5, func() error {
		// Webスクレイピングで ASINを取得する
		doc, err = goquery.NewDocument(url)
		if err != nil {
			return err
		}
		if doc.Find("title").Text() == "Amazon CAPTCHA" {
			return errors.New("Amazon CAPTCHA")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	for i := 1; i <= 20; i++ {
		var hoge string
		if i <= 3 {
			hoge = fmt.Sprintf("#zg_critical > div:nth-child(%s) >"+
				" div.a-fixed-left-grid.p13n-asin > div > div.a-fixed-left-grid-col.a-col-right > a", strconv.Itoa(i))
		} else {
			hoge = fmt.Sprintf("#zg_nonCritical > div:nth-child(%s) >"+
				" div.a-fixed-left-grid.p13n-asin > div > div.a-fixed-left-grid-col.a-col-right > a", strconv.Itoa(i-3))
		}
		doc.Find(hoge).Each(func(_ int, s *goquery.Selection) {
			url, _ := s.Attr("href")
			re := r.Copy()
			asin = append(asin, re.FindStringSubmatch(url)[0][:10])
		})
	}
	return asin, nil
}
