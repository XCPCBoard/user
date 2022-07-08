package scraper

import (
	"github.com/gocolly/colly"
)

// @Author: Feng
// @Date: 2022/4/8 17:38

//Scraper colly封装
type Scraper struct {
	C       *colly.Collector //
	threads uint32           // 启动的持久化协程数量和processor数量
}

//参数丰富接口
type scraperFunc func(*Scraper)

//WithThreads 带上协程个数
func WithThreads(threads uint32) scraperFunc {
	return func(s *Scraper) {
		s.threads = threads
	}
}

//NewScraper 构造Scraper
func NewScraper(cb func(*colly.Collector), opts ...scraperFunc) *Scraper {
	// 默认参数
	s := &Scraper{
		C: colly.NewCollector(
			colly.Async(false),
			colly.AllowURLRevisit(),
		),
		threads: 5,
	}
	// 应用外来参数
	for _, f := range opts {
		if f != nil {
			f(s)
		}
	}
	// 初始化OnHtml、OnScraped之类的
	cb(s.C)
	// 初始化
	for i := uint32(0); i < s.threads; i++ {
		// 启动持久化处理协程
		go newFlusher()
	}
	return s
}
