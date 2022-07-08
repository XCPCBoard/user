package vjudge

const (
	//24小时前的过题数
	last24HoursNumber = "vj_Person_Last_24_Hours_Pass_Number"
	//7天前的过题数
	last7DaysNumber = "vj_Person_Last_7_Days_Pass_Number"
	//30天前的过题数
	last30DaysNumber = "vj_Person_Last_30_Days_Pass_Number"
	//总过题数
	totalNumber = "vj_Person_Pass_Number"
)

//KeyWordListOfUser 用户keyWord常量列表
var KeyWordListOfUser = []string{
	last24HoursNumber, last7DaysNumber, last30DaysNumber, totalNumber,
}

//---------------------------------------------------------------------//
// 部分共用函数 //
//---------------------------------------------------------------------//

func getPersonPage(uid string) string {
	return "https://vjudge.net/user/" + uid
}
