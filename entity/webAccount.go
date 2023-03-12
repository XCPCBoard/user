package entity

import "time"

type Accounts struct {
	Id         int       `gorm:"column:id" json:"Id" form:"Id"`                         //type:int         comment:id            version:2023-02-11 21:37
	Codeforces string    `gorm:"column:codeforces" json:"Codeforces" form:"Codeforces"` //type:string      comment:codeforces    version:2023-02-11 21:37
	Nowcoder   string    `gorm:"column:nowcoder" json:"Nowcoder" form:"Nowcoder"`       //type:string      comment:nowcoder      version:2023-02-11 21:37
	Luogu      string    `gorm:"column:luogu" json:"Luogu" form:"Luogu"`                //type:string      comment:luogu         version:2023-02-11 21:37
	Atcoder    string    `gorm:"column:atcoder" json:"Atcoder" form:"Atcoder"`          //type:string      comment:atcoder       version:2023-02-11 21:37
	Vjudge     string    `gorm:"column:vjudge" json:"Vjudge" form:"Vjudge"`             //type:string      comment:vjudge        version:2023-02-11 21:37
	Rank       int       `gorm:"column:rank" json:"Rank" form:"Rank"`                   //type:int         comment:排名          version:2023-02-11 21:37
	Total      int       `gorm:"column:total" json:"Total" form:"Total"`                //type:int         comment:总过题数      version:2023-02-11 21:37
	CreatedAt  time.Time `gorm:"column:created_at" json:"CreatedAt" form:"CreatedAt"`   //type:time.Time   comment:创建时间      version:2023-02-11 21:37
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"UpdatedAt" form:"UpdatedAt"`   //type:time.Time   comment:更新时间      version:2023-02-11 21:37
}

// TableName 锁定表明
func (Accounts) TableName() string {
	return AccountTableName
}
