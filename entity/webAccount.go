package entity

import "time"

const AccountTableName = "website_account"

type Account struct {
	Id int //用户id

	CodeForces string    `gorm:"column:codeforces"` //cf
	NowCoder   string    `gorm:"column:nowcoder"`   //牛客
	LuoGu      string    `gorm:"column:luogu"`      //洛谷
	AtCoder    string    `gorm:"column:atcoder"`    //atCoder
	VJudge     string    `gorm:"column:vjudge"`     //VJ
	CreatedAt  time.Time //创建时间
	UpdatedAt  time.Time //更新时间
}

//TableName 锁定表明
func (Account) TableName() string {
	return AccountTableName
}
