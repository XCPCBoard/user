package codeforces

import (
	"XCPCer_board/scraper"
	"fmt"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

var (
	problemScraper = scraper.NewScraper(
		problemCallback,
	)
)

func problemCallback(c *colly.Collector) {
	c.OnHTML("div[style=\"position: relative;\"] #pageContent ._UserActivityFrame_frame "+
		".roundbox.userActivityRoundBox ._UserActivityFrame_footer ._UserActivityFrame_countersRow",
		func(e *colly.HTMLElement) {
			// 最近一个月过题数
			num, err := strconv.Atoi(strings.Split(e.DOM.Find(fmt.Sprintf("._UserActivityFrame_counter:contains(solved):contains("+
				"%v) ._UserActivityFrame_counterValue", lastMonthPassKeyWord)).First().Text(), " ")[0])
			if err != nil {
				log.Errorf("str atoi Error %v", err)
			} else {
				e.Request.Ctx.Put(lastMonthPassAmount, num)
			}
			// 总过题数
			num, err = strconv.Atoi(strings.Split(e.DOM.Find(fmt.Sprintf("._UserActivityFrame_counter:contains(solved):contains("+
				"%v) ._UserActivityFrame_counterValue", problemPassKeyWord)).First().Text(), " ")[0])
			if err != nil {
				log.Errorf("str atoi Error %v", err)
			} else {
				e.Request.Ctx.Put(problemPassAmountKey, num)
			}
		},
	)
}

//fetchAcceptInfo 获取过题情况
func fetchAcceptInfo(uid string) ([]scraper.KV, error) {
	// 构造上下文，及传入参数
	ctx := colly.NewContext()
	ctx.Put("uid", uid)
	// 请求
	err := problemScraper.C.Request("GET", getPersonPageUrl(uid), nil, ctx, nil)
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
