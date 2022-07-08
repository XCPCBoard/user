package vjudge

import (
	"XCPCer_board/scraper"
)

//ScrapeUser 获得所有结果
func ScrapeUser(uid string) (map[string]int, error) {
	// 请求所有并合并所有
	res, err := scraper.MergeAllResults[string, int](
		GetUserMsg(uid),
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}
