package entity

import "time"

type Post struct {
	Id        int       `gorm:"primaryKey;column:id" form:"id" json:"id" binding:"required"`   //type:int         comment:主键id              version:2023-00-01 16:58
	UserId    int       `gorm:"column:user_id" form:"userid" json:"userid" binding:"required"` //type:string       comment:用户id              version:2023-00-01 16:58
	Title     string    `gorm:"column:title" form:"title" json:"title"`                        //type:string       comment:帖子标题            version:2023-00-01 16:58
	Content   string    `gorm:"column:content"  form:"content" json:"content"`                 //type:string       comment:帖子内容            version:2023-00-01 16:58
	Note      string    `gorm:"column:note"  form:"note" json:"note"`                          //type:string       comment:备注（备用字段）    version:2023-00-01 16:58
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`                           //type:*time.Time   comment:创建时间            version:2023-00-01 16:58
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`                           //type:*time.Time   comment:更新时间            version:2023-00-01 16:58
}

// TableName 表名:posts，帖子表。
func (Post) TableName() string {
	return PostTableName
}
