package controller

import (
	log "github.com/sirupsen/logrus"
	"strconv"
	"user/entity"
	"user/service"
)

//CreatePostController 创建评论
//@param post 评论结构体指针
func CreatePostController(userId string, title string, content string, note string) error {
	uid, err := strconv.Atoi(userId)
	if err != nil {
		log.Errorf("userId is not int type :%v", err.Error())
		return err
	}
	post := entity.Post{
		UserId:  uid,
		Title:   title,
		Content: content,
		Note:    note,
	}
	return service.CreatePostService(&post)
}

//DeletePostController 删除评论
//@param id 评论id
func DeletePostController(id string) error {

	return service.DeletePostService(id)
}

//UpdatePostController 更新评论
//@param post 评论参数
func UpdatePostController(post map[string]interface{}) error {
	return service.UpdatePostService(post)
}

//SelectPostController 查询post
//@param id 评论id
//@param post 存放post的实体map
func SelectPostController(id string) (map[string]interface{}, error) {
	res := map[string]interface{}{}
	err := service.SelectPostService(id, &res)
	return res, err
}
