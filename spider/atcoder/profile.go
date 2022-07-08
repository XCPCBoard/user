package atcoder

import (
	"XCPCer_board/scraper"
	"fmt"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"strconv"
)

//---------------------------------------------------------------------//
// atCoder个人信息 //
//---------------------------------------------------------------------//
//  Key

var (
	mainScraper = scraper.NewScraper(
		mainCallback,
	)
)

//mainCallback 处理个人主页的回调函数
func mainCallback(c *colly.Collector) {
	c.OnHTML("table[class=\"dl-table mt-2\"] tbody",
		func(e *colly.HTMLElement) {
			uid := e.Request.Ctx.Get("uid")
			// 获取rating
			retRating := e.DOM.Find(fmt.Sprintf("tr:nth-child(2) span:first-child")).First().Text()
			if num, err := strconv.Atoi(retRating); err == nil {
				e.Request.Ctx.Put(getRatingKey(uid), num)
			}
			// 获取Rank
			retRank := e.DOM.Find(fmt.Sprintf("tr:nth-child(1) td")).First().Text()
			retRank = retRank[:len(retRank)-2]
			if num, err := strconv.Atoi(retRank); err == nil {
				e.Request.Ctx.Put(getRankKey(uid), num)
			}
			// 获取rating比赛场数
			retConSum := e.DOM.Find(fmt.Sprintf("tr:nth-child(4) td")).First().Text()
			if num, err := strconv.Atoi(retConSum); err == nil {
				e.Request.Ctx.Put(getContestSumKey(uid), num)
			}
		},
	)
}

//-------------------------------------------------------------------------------------------//
// 对外暴露函数
//-------------------------------------------------------------------------------------------//

//fetchMainPage 抓取个人主页页面所有
func fetchMainPage(uid string) ([]scraper.KV, error) {
	// 构造上下文，及传入参数
	ctx := colly.NewContext()
	ctx.Put("uid", uid)
	// 请求
	err := mainScraper.C.Request("GET", getAtCoderBaseUrl(uid), nil, ctx, nil)
	if err != nil {
		log.Errorf("scraper error %v", err)
		return nil, err
	}
	// 解构出kv对
	kvs := scraper.Parse(ctx, map[string]struct{}{
		"uid": {},
	})
	return kvs, nil
}
