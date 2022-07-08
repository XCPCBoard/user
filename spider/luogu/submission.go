package luogu

import (
	"XCPCer_board/scraper"
	"encoding/json"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"net/url"
	"strconv"
	"strings"
)

//---------------------------------------------------------------------//
// 获取int形式的信息 //
//---------------------------------------------------------------------//

var (
	subScraper *scraper.Scraper[ProblemPass]
)

func init() {
	subScraper = scraper.NewScraper[ProblemPass](
		scraper.WithCallback(subCallback),
		scraper.WithThreads[ProblemPass](2),
	)
}

func subCallback(c *colly.Collector, res *scraper.Results[ProblemPass]) {
	c.OnHTML("head", func(e *colly.HTMLElement) {

		//decoder
		text, _ := url.QueryUnescape(e.DOM.Text())
		//get JsonText
		Data := text[strings.Index(text, "{") : strings.LastIndex(text, "}")+1]
		err := json.Unmarshal([]byte(Data), &jsonData)
		if err != nil {
			log.Println("json Unmarshal error: ", err)
		}
		if jsonData.Code != 200 {
			log.Println("http Response is not 200: ", err)
		}

		problem := jsonData.GetCurrentData().GetPassedProblems()
		for _, i := range problem {
			tp := ProblemPass{
				dif:   strconv.Itoa(int(i.GetDifficulty())),
				title: i.GetTitle(),
			}
			res.Set(i.GetPid(), tp)
		}
	})
}

func GetSubMsg(uid string) scraper.Results[ProblemPass] {
	return subScraper.Scrape(getPersonPractice(uid))
}

type ProblemPass struct {
	//题名
	title string
	//难度
	dif string
}
