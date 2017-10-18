package main


import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var myDB *sql.DB

func DataBaseInit(){
	dataSource := os.Getenv("DATABASE_URL")
	var err error
	myDB, err = sql.Open("mysql", dataSource) //"root:@/my_database")
	if err != nil {
		panic(err)
	}
}

// get rank urls
// WANG this is 10 time second. only go func{}()
func getURL()([]string, error){
	rows, err := myDB.Query("SELECT URL FROM CategoryURL;")
	if err != nil {
		return nil, err
	}
	// list append
	var urls []string = make([]string,0)
	for rows.Next() {
		var url string
		if err := rows.Scan(&url); err != nil{
			return nil, err
		}
		urls = append(urls, url)
	}
	return urls,nil
}

func main() {

	DataBaseInit()
	getURL()
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
