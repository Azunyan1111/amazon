package main

import (
	"github.com/Azunyan1111/amazon/model"
	"time"
	"fmt"
)

func main() {
	model.DataBaseInit()

	start := time.Now()
	model.GetRankingASIN()
	//log.Println(model.GetUrl())
	end := time.Now()
	fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())
}
