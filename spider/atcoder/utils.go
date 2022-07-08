package atcoder

import (
	"fmt"
	"strconv"
)

const (
	RatingKey     = "atc_rating"
	contestSumKey = "atc_contest_sum"
	rankKey       = "atc_rank"
	submissionKey = "atc_submission"
)

func getRatingKey(uid string) string {
	return fmt.Sprintf("%v_%v", RatingKey, uid)
}

func getRankKey(uid string) string {
	return fmt.Sprintf("%v_%v", rankKey, uid)
}

func getContestSumKey(uid string) string {
	return fmt.Sprintf("%v_%v", contestSumKey, uid)
}

func getSubmissionKey(uid string, cid string, pid string) string {
	return fmt.Sprintf("%v_%v_%v_%v", submissionKey, uid, cid, pid)
}

//getPageUrl 获取比赛列表url
func getPageUrl(page int) string {
	return "https://atcoder.jp/contests/archive?page=" + strconv.Itoa(page)
}

//getSubmissionPageUrl 获取用户提交界面url
func getSubmissionPageUrl(cid string, uid string) string {
	return "https://atcoder.jp/contests/" + cid + "/submissions?f.User=" + uid + "&f.Status=AC"
}

//getAtCoderBaseUrl 获取个人主页URL
func getAtCoderBaseUrl(atCoderId string) string {
	return "https://atcoder.jp/users/" + atCoderId
}
