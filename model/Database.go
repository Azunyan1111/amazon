package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var myDB *sql.DB

func DataBaseInit() {
	//hoge := "root:541279xx@tcp(mydbinstance.cv8ap3ddulzc.us-east-2.rds.amazonaws.com:3306)/amazon"
	dataSource := os.Getenv("DATABASE_URL")
	var err error
	myDB, err = sql.Open("mysql", dataSource) //"root:@/my_database")
	if err != nil {
		panic(err)
	}
}

// get rank urls
// WANG this is 10 time second. only go func{}()
func GetUrl() ([]string, error) {
	// TODO: LIMIT
	rows, err := myDB.Query("SELECT URL FROM CategoryURL;")
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
		_, err := myDB.Exec("INSERT INTO ASIN(ASIN) VALUES(?)",asin)
		if err != nil {
			continue
		}
	}
}

func mains() {

	DataBaseInit()
	GetUrl()
	/*//
	rows, err := myDB.Query("SELECT id,URL FROM CategoryURL;")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	type test struct {
		Id       int64
		Category string
	}

	start := time.Now()
	var posts []test
	for rows.Next() {
		var post test
		if err := rows.Scan(&post.Id, &post.Category); err != nil{
			panic(err)
		}
		posts = append(posts, post)
	}
	// 処理
	end := time.Now()
	fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())
	//fmt.Println(posts)

	//*/
	// データベース接続終了(リターン時呼び出し)
	defer myDB.Close()
}
