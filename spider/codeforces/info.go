package codeforces

import (
	"XCPCer_board/model"
	"XCPCer_board/scraper"
	"encoding/json"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
)

// @Author: Feng
// @Date: 2022/5/12 22:41

var (
	infoScraper = scraper.NewScraper(
		userInfoCallback,
	)
)

//userInfoCallback 处理codeforces的api
func userInfoCallback(c *colly.Collector) {
	c.OnScraped(func(r *colly.Response) {
		// 获取uid
		uid := r.Request.Ctx.Get("uid")
		if uid == "" {
			log.Errorf("%v", model.UidError)
			return
		}
		// 反序列化
		rsp := &UserInfo{}
		err := json.Unmarshal(r.Body, rsp)
		if err != nil {
			log.Errorf("Codeforces User Info Unmarshal Error %v", err)
			return
		}
		if rsp.GetStatus() != "OK" || len(rsp.GetResult()) != 1 {
			log.Errorf("Response: %v Infos Length: %v", rsp.GetStatus(), len(rsp.GetResult()))
			return
		}
		info := rsp.GetResult()[0]
		if info.GetRating() != 0 {
			r.Ctx.Put(GetRatingKey(uid), info.GetRating())
		}
		if info.GetMaxRating() != 0 {
			r.Ctx.Put(GetMaxRatingKey(uid), info.GetMaxRating())
		}
		if info.GetRank() != "" {
			r.Ctx.Put(GetRankingNameKey(uid), info.GetRank())
		}
		if info.GetMaxRank() != "" {
			r.Ctx.Put(GetMaxRankingNameKey(uid), info.GetMaxRank())
		}
	})
}

//---------------------------------------------------------------------//
// 对外暴露函数:用户信息获取
//---------------------------------------------------------------------//

//fetchUserInfo 抓取用户信息
func fetchUserInfo(uid string) ([]scraper.KV, error) {
	// 构造上下文，及传入参数
	ctx := colly.NewContext()
	ctx.Put("uid", uid)
	// 请求
	err := infoScraper.C.Request("GET", getUserInfoUrl(uid), nil, ctx, nil)
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
