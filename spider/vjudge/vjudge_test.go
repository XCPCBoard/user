package vjudge

import (
	_ "XCPCer_board/config"
	_ "XCPCer_board/dao"
	"XCPCer_board/model"
	"testing"
)

////////////////////////////////////////////////////
///////////////////  测试用例  //////////////////////
////////////////////////////////////////////////////
const (
	testPackage = "vJudge"
)

type userExample struct {
	msg map[string]int
}

//userSet 用户集合
var userSet = map[string]userExample{
	model.TestVJIdLYF: userExample{
		msg: map[string]int{
			last24HoursNumber: 0,
			last7DaysNumber:   0,
			last30DaysNumber:  0,
			totalNumber:       30,
		},
	},
}

////////////////////////////////////////////////////
///////////////////  主测试函数  /////////////////////
////////////////////////////////////////////////////

//UserMsgTest 用户信息测试函数
func UserMsgTest(t *testing.T) {

	//test ScrapeUser
	for uid, correctMsg := range userSet {
		//get msg
		funcRet, err := ScrapeUser(uid)
		if err != nil {
			t.Errorf("Errorin all msg: %v", err)
		}
		//check map
		if len(correctMsg.msg) != len(funcRet) {
			t.Errorf("Errorin all msg\n ret= %v  \nbut the ans is %v", funcRet, correctMsg.msg)
		}
		for k, v := range correctMsg.msg {
			if _, r := funcRet[k]; r == false || v != funcRet[k] {
				t.Errorf("Error in all msg\n ret= %v  \nbut the ans is %v", funcRet, correctMsg.msg)
			}
		}
	}

}

//TestLg luoGu总测试函数
func TestLg(t *testing.T) {
	UserMsgTest(t)
	//SubmissionTest(t)
}
