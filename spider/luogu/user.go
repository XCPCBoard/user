package luogu

import (
	"XCPCer_board/scraper"
	"encoding/json"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"net/url"
	"strings"
)

//---------------------------------------------------------------------//
// 获取int形式的信息 //
//---------------------------------------------------------------------//

var (
	intScraper *scraper.Scraper[int]
	jsonData   UserShow
	difficulty [5]int
)

func init() {
	intScraper = scraper.NewScraper[int](
		scraper.WithCallback(intCallback),
		scraper.WithThreads[int](2),
	)
}

func intCallback(c *colly.Collector, res *scraper.Results[int]) {
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

		//count problem difficulty
		user := jsonData.GetCurrentData().GetUser()
		problem := jsonData.GetCurrentData().GetPassedProblems()

		for i := 0; i < 5; i++ {
			difficulty[i] = 0
		}
		for _, i := range problem {
			q := i.GetDifficulty()
			if q == 0 || q > 7 { //未知题
				difficulty[0]++
			} else if q < 2 { //入门就是简单 q=1
				difficulty[1]++
			} else if q < 4 { //普及-就是基础 q=2,3
				difficulty[2]++
			} else if q < 6 { //普及/提高-,普及+/提高 是进阶 q=4,5
				difficulty[3]++
			} else { //困难
				difficulty[4]++
			}
		}

		//set data
		res.Set(passProblemNumber, int(user.GetPassedProblemCount()))
		res.Set(ranting, int(user.GetRanking()))
		//set data of problem
		res.Set(unKnowProblem, difficulty[0])
		res.Set(simpleProblem, difficulty[1])
		res.Set(basicProblem, difficulty[2])
		res.Set(elevatedProblem, difficulty[3])
		res.Set(hardProblem, difficulty[4])
	})
}

//////////////////////////////////////
///////////    对外暴露     ///////////
//////////////////////////////////////

// GetUserMsg  获取用户信息
func GetUserMsg(uid string) scraper.Results[int] {
	return intScraper.Scrape(getPersonPractice(uid))
}

//UserMsg 用户信息结构体 ，暂时无用
type UserMsg struct {
	Uid               string
	PassProblemNumber int
	Ranting           int
	SimpleProblem     int
	BasicProblem      int
	ElevatedProblem   int
	HardProblem       int
	UnKnowProblem     int
}

//StructToMap 结构体转Map，暂时无用
func StructToMap(user UserMsg) (map[string]int, string) {
	var mp map[string]int
	mp[passProblemNumber] = user.PassProblemNumber
	mp[ranting] = user.Ranting
	mp[simpleProblem] = user.SimpleProblem
	mp[basicProblem] = user.BasicProblem
	mp[elevatedProblem] = user.ElevatedProblem
	mp[hardProblem] = user.HardProblem
	mp[unKnowProblem] = user.UnKnowProblem
	return mp, user.Uid
}

//MapToStruct Map转结构体, 返回的bool=1为正常，0为map里没有该值，暂时无用
func MapToStruct(mp map[string]int) (UserMsg, bool) {

	var user UserMsg
	var ok bool

	if user.PassProblemNumber, ok = mp[passProblemNumber]; !ok {
		return user, ok
	}
	if user.Ranting, ok = mp[ranting]; !ok {
		return user, ok
	}
	if user.SimpleProblem, ok = mp[simpleProblem]; !ok {
		return user, ok
	}
	if user.BasicProblem, ok = mp[basicProblem]; !ok {
		return user, ok
	}
	if user.ElevatedProblem, ok = mp[elevatedProblem]; !ok {
		return user, ok
	}
	if user.HardProblem, ok = mp[hardProblem]; !ok {
		return user, ok
	}
	if user.UnKnowProblem, ok = mp[unKnowProblem]; !ok {
		return user, ok
	}
	return user, true
}
