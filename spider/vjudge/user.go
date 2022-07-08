package vjudge

import (
	"XCPCer_board/scraper"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"strconv"
)

//---------------------------------------------------------------------//
// 获取int形式的信息 //
//---------------------------------------------------------------------//

var (
	intScraper *scraper.Scraper[int]
)

func init() {
	intScraper = scraper.NewScraper[int](
		scraper.WithCallback(intCallback),
		scraper.WithThreads[int](2),
	)
}

func intCallback(c *colly.Collector, res *scraper.Results[int]) {
	c.OnHTML("body", func(e *colly.HTMLElement) {
		//fmt.Println(e.DOM.First().Text())
		res.Set(last24HoursNumber, last24HoursNumberHandler(e.DOM))
		res.Set(last7DaysNumber, last7DaysNumberHandler(e.DOM))
		res.Set(last30DaysNumber, last30DaysNumberHandler(e.DOM))
		res.Set(totalNumber, totalNumberHandler(e.DOM))
	})
}

//last24HoursNumberHandler 获取24小时前的过题数
func last24HoursNumberHandler(doc *goquery.Selection) int {
	retStr := doc.Find(fmt.Sprintf(".container a[title=\"New solved in last 24 hours\"]")).First().Text()
	num, err := strconv.Atoi(retStr)
	if err != nil {
		log.Errorf("VJ strToInt get err:%v\t and the return is %v:", retStr, err)
		return -1
	}
	return num
}

//last7DaysNumberHandler 获取7天前的过题数
func last7DaysNumberHandler(doc *goquery.Selection) int {
	retStr := doc.Find(fmt.Sprintf(".container a[title=\"New solved in last 7 days\"]")).First().Text()
	num, err := strconv.Atoi(retStr)
	if err != nil {
		log.Errorf("VJ strToInt get err:%v\t and the return is %v:", retStr, err)
		return -1
	}
	return num
}

//last30DaysNumberHandler 获取一个月前的过题数
func last30DaysNumberHandler(doc *goquery.Selection) int {
	retStr := doc.Find(fmt.Sprintf(".container a[title=\"New solved in last 30 days\"]")).First().Text()
	num, err := strconv.Atoi(retStr)
	if err != nil {
		log.Errorf("VJ strToInt get err:%v\t and the return is %v:", retStr, err)
		return -1
	}
	return num
}

//totalNumberHandler 获取总的过题数
func totalNumberHandler(doc *goquery.Selection) int {
	retStr := doc.Find(fmt.Sprintf(".container a[title=\"Overall solved\"]")).First().Text()
	num, err := strconv.Atoi(retStr)
	if err != nil {
		log.Errorf("VJ strToInt get err:%v\t and the return is %v:", retStr, err)
		return -1
	}
	return num
}

//////////////////////对外暴露函数////////////////////////

//GetUserMsg 获取用户信息
func GetUserMsg(uid string) scraper.Results[int] {
	return intScraper.Scrape(getPersonPage(uid))
}
