package entity

import "time"

type Posts struct {
	Id        int       `gorm:"column:id" json:"Id" form:"Id"`                       //type:int         comment:主键id              version:2023-02-05 21:39
	UserId    int       `gorm:"column:user_id" json:"UserId" form:"UserId"`          //type:int         comment:用户id              version:2023-02-05 21:39
	Title     string    `gorm:"column:title" json:"Title" form:"Title"`              //type:string      comment:帖子标题            version:2023-02-05 21:39
	Content   string    `gorm:"column:content" json:"Content" form:"Content"`        //type:string      comment:帖子内容            version:2023-02-05 21:39
	Note      string    `gorm:"column:note" json:"Note" form:"Note"`                 //type:string      comment:备注（备用字段）    version:2023-02-05 21:39
	CreatedAt time.Time `gorm:"column:created_at" json:"CreatedAt" form:"CreatedAt"` //type:time.Time   comment:创建时间            version:2023-02-05 21:39
	UpdatedAt time.Time `gorm:"column:updated_at" json:"UpdatedAt" form:"UpdatedAt"` //type:time.Time   comment:更新时间            version:2023-02-05 21:39
}

// TableName 表名:posts，帖子表。
func (Posts) TableName() string {
	return PostTableName
}
