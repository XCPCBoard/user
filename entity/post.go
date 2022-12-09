package entity

import "time"

type Post struct {
	Id int //评论id

	UserId  int    //用户id
	Title   string //帖子标题
	Content string //帖子内容
	Note    string //备注（预留字段)

	CreatedAt time.Time //创建时间
	UpdatedAt time.Time //更新时间
}
