package entity

type Post struct {
	Id int //评论id

	userId    int    //用户id
	Title     string //帖子标题
	content   string //帖子内容
	note      string //备注（预留字段)
	CreatedAt string //创建时间
	UpdatedAt string //更新时间
}
