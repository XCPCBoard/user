package luogu

const (
	////////////////user//////////////////
	//过题数
	passProblemNumber = "problem_number"
	//排名
	ranting = "ranting"
	//简单题个数
	simpleProblem = "simple_problem_number"
	//基础题个数
	basicProblem = "base_problem_number"
	//提高题个数
	elevatedProblem = "elevated_problem_number"
	//困难题个数
	hardProblem = "hard_problem_number"
	//未知题个数
	unKnowProblem = "unKnow_problem_number"
)

//KeyWordListOfUser 用户keyWord常量列表
var KeyWordListOfUser = []string{passProblemNumber, ranting, simpleProblem,
	basicProblem, elevatedProblem, hardProblem, unKnowProblem}

//获取网页函数
func getPersonPage(uid string) string {
	return "https://www.luogu.com.cn/user/" + uid
}
func getPersonPractice(uid string) string {
	return getPersonPage(uid) + "#practice"
}
