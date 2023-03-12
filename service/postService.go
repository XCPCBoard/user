package service

import (
	"fmt"
	"github.com/XCPCBoard/common/dao"
	"github.com/XCPCBoard/common/errors"
	"github.com/XCPCBoard/common/logger"
	"github.com/XCPCBoard/user/entity"
)

// CreatePostService 创建评论
// @param post 评论结构体指针
func CreatePostService(post *entity.Posts) *errors.MyError {
	res := dao.DBClient.Create(post)
	if res.Error != nil {
		logger.L.Error("create post error", res.Error, 0, fmt.Sprintf("%+v", post))
		return errors.GetError(errors.INNER_ERROR, "创建评论异常")
	}
	return nil
}

// DeletePostService 删除评论
// @param id 评论id
func DeletePostService(id string) *errors.MyError {

	res := dao.DBClient.Delete(&entity.Posts{}, id)
	return CreatError(res, "DeletePostService error", fmt.Sprintf("delete post id :%v", id))
}

// UpdatePostService 更新评论
// @param post 评论参数
func UpdatePostService(post *entity.Posts) *errors.MyError {
	////检查主键是否为0,gorm特性是id为
	//if post.Id == 0 {
	//	err := errors.New("post's id should not be 0")
	//	logger.L.Error("post's id should not be 0", err,0, fmt.Sprintf("post: %+v", post))
	//	return err
	//}

	res := dao.DBClient.Model(&entity.Posts{}).Where("id = ?", post.Id).Updates(post)
	return CreatError(res, "Update post service error", fmt.Sprintf("post: %+v", post))
}

// SelectPostService 查询post
// @param id 评论id
// @param post 存放post的实体map
func SelectPostService(id string, post *entity.Posts) *errors.MyError {

	res := dao.DBClient.Model(&entity.Posts{}).Find(post, id)
	return CreatError(res, "SelectPostService error", fmt.Sprintf("can not find post:%v", id))
}
