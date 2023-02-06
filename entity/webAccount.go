package entity

import "time"

type Account struct {
	Id int `gorm:"primaryKey;column:id" form:"id" binding:"required" ` //用户id

	CodeForces string    `gorm:"column:codeforces" form:"codeforces" ` //cf
	NowCoder   string    `gorm:"column:nowcoder" form:"nowcoder" `     //牛客
	LuoGu      string    `gorm:"column:luogu" form:"luogu" `           //洛谷
	AtCoder    string    `gorm:"column:atcoder" form:"atcoder" `       //atCoder
	VJudge     string    `gorm:"column:vjudge" form:"vjudge" `         //VJ
	CreatedAt  time.Time `gorm:"column:created_at" json:"createtime"`  //type:*time.Time   comment:创建时间
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updatetime"`  //type:*time.Time   comment:更新时间

}

// TableName 锁定表明
func (Account) TableName() string {
	return AccountTableName
}
