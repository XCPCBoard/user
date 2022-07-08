package luogu

import (
	"XCPCer_board/scraper"
)

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
func ScrapeSub(uid string) (map[string]ProblemPass, error) {
	res, err := scraper.MergeAllResults[string, ProblemPass](
		GetSubMsg(uid),
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}
