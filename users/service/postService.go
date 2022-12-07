package service

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"user/dao"
	"user/users/entity"
	"user/users/util"
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
	return util.CreatError(res, fmt.Sprintf("can not find post:%v", id))
}

//UpdatePostService 更新评论
//@param post 评论参数
func UpdatePostService(post map[string]interface{}) error {
	//检查是否包含主键
	if _, ok := post["id"]; !ok {
		err := errors.New("can't find post's id")
		log.Errorf(err.Error())
		return err
	}

	res := dao.DBClient.Model(&entity.Post{}).Where("id = ?", post["id"]).Updates(post)
	return util.CreatError(res, fmt.Sprintf("can not find post:%v", post["id"]))
}

func SelectPostService(id string, post *map[string]interface{}) error {

	res := dao.DBClient.Model(&entity.Post{}).Find(post, id)
	return util.CreatError(res, fmt.Sprintf("can not find post:%v", id))
}
