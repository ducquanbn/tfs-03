package crawlers

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	entity "github.com/ducquan/Crawler/entity"
	"gorm.io/gorm"
)

func CrawlerFilm(url string, db *gorm.DB) {
	fimls := make(map[int]entity.Film)
	res, _ := http.Get(url)
	data, _ := goquery.NewDocumentFromReader(res.Body)
	data.Find(".lister-list tr").Each(func(i int, s *goquery.Selection) {
		temp := entity.Film{}
		temp.Rank = i
		temp.Rating, _ = s.Find(".posterColumn [name = ir]").Attr("data-value")
		temp.Year = s.Find(".titleColumn span").Text()
		info := s.Find(".titleColumn a")
		temp.Name = info.Text()
		temp.Link, _ = info.Attr("href")
		temp.Link = "" + temp.Link
		title, _ := info.Attr("title")
		titleSplit := strings.Split(title, ", ")
		temp.Director = titleSplit[0]
		temp.Writers = strings.Join(titleSplit[1:], ", ")
		fimls[i+1] = temp
	})
	db.AutoMigrate(&entity.Film{})
	for i := 1; i <= len(fimls); i++ {
		db.Create(&entity.Film{Rank: fimls[i].Rank, Name: fimls[i].Name, Year: fimls[i].Year, Rating: fimls[i].Rating, Director: fimls[i].Director, Link: fimls[i].Link, Writers: fimls[i].Writers})
	}
}
