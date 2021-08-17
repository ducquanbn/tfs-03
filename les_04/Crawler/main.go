package main

import (
	"fmt"

	craw "github.com/ducquan/Crawler/crawlers"
	db "github.com/ducquan/Crawler/database"
)

func main() {
	con := db.DBConnect()
	url := "https://www.imdb.com/chart/top/?ref_=nv_mv_250"
	craw.CrawlerFilm(url, con)
	fmt.Println("Crawler thanh cong!!!!!!")

}
