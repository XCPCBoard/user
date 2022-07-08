package user

//用户表
//用户id（主键，自动生成) 账号	密码(账号+用户密码的哈希存储）
// 	姓名		各个网站id（必填）	管理员标签(bool，默认0，只有管理员能修改为1)	邮箱（必填，用于恢复账号密码）
//头像（有默认头像)	个签（默认空）		手机（选填，默认空)		QQ(选填,默认空)

//user 用户变量
type user struct {
	id int //用户id

	account         string //账号
	keyword         string //密码
	email           string //邮箱
	isAdministrator bool   //管理员标签(=1时为管理员)

	name        string      //姓名
	listOfWeb   webSiteList //网站列表
	imagePath   string      //头像的路径
	signature   string      //个性签名
	phoneNumber string      //手机号
	QQNumber    string      //qq号

}

//网站列表
type webSiteList struct {
	codeForces string //cf
	nowCoder   string //牛客
	luoGu      string //洛谷
	atCoder    string //atCoder
	vJudge     string //VJ
}
