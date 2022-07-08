package nowcoder

import (
	"XCPCer_board/scraper"
	log "github.com/sirupsen/logrus"
)

// @Author: Feng
// @Date: 2022/4/8 17:09

var (
	// 爬取函数
	fetchers = []func(uid string) ([]scraper.KV, error){
		fetchMainPage,
		fetchPractice,
	}
)

//scrape 拉取牛客的所有结果
func scrape(uid string) (res []scraper.KV) {
	// 请求所有
	for _, f := range fetchers {
		// 请求
		kvs, err := f(uid)
		if err != nil {
			log.Errorf("GetPersistHandler Fetcher Error %v", err)
			continue
		}
		res = append(res, kvs...)
	}
	return res
}

//Flush 刷新某用户牛客id信息
func Flush(uid string) {
	// 拉出所有kv对
	kvs := scrape(uid)
	// 向持久化处理协程注册持久化处理函数
	scraper.CustomFlush(func() error {
		log.Infoln(kvs)
		return nil
	})
}
