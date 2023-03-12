package handle

import (
	"fmt"
	response "github.com/XCPCBoard/api/http"
	"github.com/XCPCBoard/common/errors"
	"github.com/XCPCBoard/common/logger"
	"github.com/XCPCBoard/user/entity"
	"github.com/XCPCBoard/user/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

// CreatePostHandle 创建评论,只能自己创建自己的评论
func CreatePostHandle(ctx *gin.Context) {

	userId := ctx.PostForm("userId")
	title := ctx.PostForm("title")
	content := ctx.PostForm("content")
	note := ctx.PostForm("note")

	//对比tokenID和userID
	errToken := CheckUserIDIsCorrect(ctx, userId)
	if errToken != nil {
		return
	}
	//转为int
	uid, err := strconv.Atoi(userId)
	if err != nil {
		e := errors.CreateError(errors.ERROR.Code, "uid转为Int失败", userId)
		logger.L.Err(e, 0)
		ctx.Error(e)
		return
	}

	post := entity.Posts{
		UserId:  uid,
		Title:   title,
		Content: content,
		Note:    note,
	}
	//业务
	if err := service.CreatePostService(&post); err != nil {
		ctx.Error(err)
		return
	}

	response.SuccessResponseAddToken(ctx, nil)
}

// DeleteSelfPostHandle 删除评论
func DeleteSelfPostHandle(ctx *gin.Context) {
	//获取参数
	id := ctx.PostForm("id")
	_logTheCtx := fmt.Sprintf("%#v , %#v", ctx.Params, ctx.Keys)

	//获取评论对应的id
	post := new(entity.Posts)
	if err := service.SelectPostService(id, post); err != nil {
		logger.L.Err(err, 0)
		ctx.Error(err)
		return
	}

	//对比tokenID和userID
	errToken := CheckUserIDIsCorrect(ctx, fmt.Sprintf("%v", post.UserId))
	if errToken != nil {
		return
	}
	//删除
	if err := service.DeletePostService(id); err != nil {
		logger.L.Error("删除评论错误", err, 0, _logTheCtx)
		ctx.Error(err)
		return
	}

	response.SuccessResponseAddToken(ctx, nil)

}

// UpdateSelfPostHandle 更新评论
func UpdateSelfPostHandle(ctx *gin.Context) {

	//获取参数
	data := &entity.Posts{}
	_logTheCtx := fmt.Sprintf("%#v , %#v", ctx.Params, ctx.Keys)

	if err := ctx.ShouldBind(data); err != nil {
		e := errors.CreateError(errors.ERROR.Code, "获取update参数错误", _logTheCtx)
		logger.L.Err(e, 0)
		ctx.Error(e)
		return
	}

	//检查id和token中的Id是否相符
	errToken := CheckUserIDIsCorrect(ctx, strconv.Itoa(data.UserId))
	if errToken != nil {
		return
	}
	//更新
	if err := service.UpdatePostService(data); err != nil {
		ctx.Error(err)
		return
	}

	response.SuccessResponseAddToken(ctx, nil)

}

// SelectPostHandle 查询post
func SelectPostHandle(ctx *gin.Context) {
	//获取参数
	id := ctx.Query("id")

	//无检测，任何有合法token的人都可以看到任意post

	//查询
	res := new(entity.Posts)
	if err := service.SelectPostService(id, res); err != nil {
		ctx.Error(err)
		return
	}

	m := make(map[string]interface{})
	if err := StrToMap(ctx, res, &m); err != nil {
		logger.L.Err(err, 0)
		return
	}
	response.SuccessResponseAddToken(ctx, m)

}
