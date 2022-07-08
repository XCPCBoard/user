package codeforces

import "fmt"

//---------------------------------------------------------------------//
// 常量
//---------------------------------------------------------------------//

const (
	// 个人rating
	ratingKey = "codeforces_rating"
	// 个人历史最高rating
	maxRatingKey = "codeforces_max_rating"
	//当前rating所对应的等级（红名、紫名...)
	rankingNameKey = "codeforces_ranking_name"
	//最大rating所对应的等级（红名、紫名...)
	maxRankingNameKey = "codeforces_max_ranking_name"
	// 个人总过题数
	problemPassAmountKey = "codeforces_problem_pass"
	// 个人最后一月过题数
	lastMonthPassAmount = "codeforces_last_month_problem_pass"

	// CF finder关键词
	// 个人总过题数
	problemPassKeyWord = "all"
	//个人最后一月过题数
	lastMonthPassKeyWord = "month"
)

//---------------------------------------------------------------------//
// 共用函数
//---------------------------------------------------------------------//

func getPersonPageUrl(uid string) string {
	return "https://codeforces.com/profile/" + uid
}

func getUserInfoUrl(uid string) string {
	return "https://codeforces.com/api/user.info?handles=" + uid
}

func GetRatingKey(uid string) string {
	return fmt.Sprintf("%v_%v", ratingKey, uid)
}

func GetMaxRatingKey(uid string) string {
	return fmt.Sprintf("%v_%v", maxRatingKey, uid)
}

func GetRankingNameKey(uid string) string {
	return fmt.Sprintf("%v_%v", rankingNameKey, uid)
}

func GetMaxRankingNameKey(uid string) string {
	return fmt.Sprintf("%v_%v", maxRankingNameKey, uid)
}
