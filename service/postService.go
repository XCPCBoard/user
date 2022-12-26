package service

import (
	"errors"
	"fmt"
	"github.com/XCPCBoard/user/dao"
	"github.com/XCPCBoard/user/entity"
	log "github.com/sirupsen/logrus"
)

//CreatePostService 创建评论
//@param post 评论结构体指针
func CreatePostService(post *entity.Post) error {
	res := dao.DBClient.Create(post)
	if res.Error != nil {
		log.Errorf("Create Post Service error")
		return res.Error
	}
	return nil
}

//DeletePostService 删除评论
//@param id 评论id
func DeletePostService(id string) error {

	res := dao.DBClient.Delete(&entity.Post{}, id)
	return CreatError(res, fmt.Sprintf("can not find post:%v", id))
}

//UpdatePostService 更新评论
//@param post 评论参数
func UpdatePostService(post map[string]string) error {
	//检查是否包含主键
	if _, ok := post["id"]; !ok {
		err := errors.New("can't find post's id")
		log.Errorf(err.Error())
		return err
	}

	res := dao.DBClient.Model(&entity.Post{}).Where("id = ?", post["id"]).Updates(post)
	return CreatError(res, fmt.Sprintf("can not find post:%v", post["id"]))
}

//SelectPostService 查询post
//@param id 评论id
//@param post 存放post的实体map
func SelectPostService(id string, post *map[string]interface{}) error {

	res := dao.DBClient.Model(&entity.Post{}).Find(post, id)
	return CreatError(res, fmt.Sprintf("can not find post:%v", id))
}
