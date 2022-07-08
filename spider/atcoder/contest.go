package atcoder

import (
	"XCPCer_board/model"
	"XCPCer_board/scraper"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

var (
	conScraper = scraper.NewScraper(
		conCallback,
	)
	contestId string
	maxPage   = 1 //比赛列表页数
)

// submission 信息
type submission struct {
	userid string //用户名
	CTid   string //比赛编号
	task   string //题目序号
	score  int    //题目难度
	SMid   string //提交编号
}

//conCallback 处理比赛列表的回调函数
func conCallback(c *colly.Collector) {
	d := c.Clone()
	// 获取用户比赛提交页面信息
	d.OnHTML("table[class=\"table table-bordered table-striped small th-center\"] tbody tr",
		func(e *colly.HTMLElement) {
			uid := e.Request.Ctx.Get("uid")
			if uid == "" {
				log.Errorf("%v", model.UidError)
				return
			}
			//题目序号
			task := strings.Split(e.DOM.Find("td:nth-child(2)").First().Text(), " ")[0]
			if task == "" {
				log.Errorf("task is empty")
				return
			}
			//题目难度
			score, errSc := strconv.Atoi(e.DOM.Find("td:nth-child(5)").First().Text())
			if errSc != nil {
				log.Errorf("Submission Score Fetcher Error %v", errSc)
				return
			}
			//提交编号
			SMidStr := strings.Split(e.ChildAttr("td:nth-child(10) a", "href"), "/")
			SMid := SMidStr[len(SMidStr)-1]
			if SMid == "" {
				log.Errorf("submission id is empty")
				return
			}
			e.Request.Ctx.Put(getSubmissionKey(uid, contestId, task), submission{uid, contestId, task, score, SMid})
		})

	//获取所有比赛id
	c.OnHTML("div[class=\"col-lg-9 col-md-8\"]",
		func(e *colly.HTMLElement) {
			//获取比赛列表页数
			np, err := strconv.Atoi(e.DOM.Find("div[class=\"text-center\"] ul li:last-child").First().Text())
			if err != nil {
				log.Errorf("Atcoder Page Error %v", err)
				return
			}
			maxPage = np

			uid := e.Request.Ctx.Get("uid")
			if uid == "" {
				log.Errorf("%v", model.UidError)
				return
			}
			// 访问每个页面的contest
			e.ForEach("tbody tr", func(i int, element *colly.HTMLElement) {
				//比赛id
				cLink := element.ChildAttr("td:nth-child(2) a", "href")
				contestId = strings.Split(cLink, "/")[2]
				if contestId == "" {
					log.Errorf("contest id is empty")
					return
				}
				//部分比赛无访问比赛提交记录权限
				if err := d.Request("GET", getSubmissionPageUrl(contestId, uid), nil, e.Request.Ctx, nil); err != nil {
					if err.Error() != "Not Found" {
						log.Errorf("atcoder submissionpage error : %v", err)
					}
				}
			})
		})
}

//-------------------------------------------------------------------------------------------//
// 对外暴露函数
//-------------------------------------------------------------------------------------------//

//fetchConPage 抓取用户提交所有提交信息
func fetchConPage(uid string) ([]scraper.KV, error) {
	// 构造上下文，及传入参数
	var res []scraper.KV
	for i := 1; i <= maxPage; i++ {
		ctx := colly.NewContext()
		ctx.Put("uid", uid)
		// 请求
		err := conScraper.C.Request("GET", getPageUrl(i), nil, ctx, nil)
		if err != nil {
			log.Errorf("atcoder contestpage error %v", err)
			break
		}
		// 解构出kv对
		kvs := scraper.Parse(ctx, map[string]struct{}{
			"uid": {},
		})
		res = append(res, kvs...)
	}
	return res, nil
}
