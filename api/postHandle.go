package api

import (
	"fmt"
	response "github.com/XCPCBoard/api/http"
	"github.com/XCPCBoard/common/errors"
	"github.com/XCPCBoard/common/logger"
	"github.com/XCPCBoard/user/entity"
	"github.com/XCPCBoard/user/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreatePostHandle 创建评论
func CreatePostHandle(ctx *gin.Context) {

	userId := ctx.PostForm("userId")
	title := ctx.PostForm("title")
	content := ctx.PostForm("content")
	note := ctx.PostForm("note")

	//resErr:=new(errors.MyError)

	//对比tokenID和userID
	if !CheckUserIDIsCorrect(ctx, userId) {

		logger.Logger.Error("create post error", nil, 0, fmt.Sprintf("userId :%v", userId))
		ctx.Error(errors.GetError(errors.VALID_ERROR, "用户非本人"))
		return
	}

	//转为int
	uid, err := strconv.Atoi(userId)
	if err != nil {

		logger.Logger.Error("create post error", err, 0, fmt.Sprintf("userId :%v", userId))
		ctx.Error(errors.GetError(errors.VALID_ERROR, "userId is not int type "))
		return
	}

	post := entity.Post{
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

	ctx.JSON(http.StatusOK, response.SuccessResponse("ok", nil))
}

// DeletePostHandle 删除评论
func DeletePostHandle(ctx *gin.Context) {
	//获取参数
	id := ctx.PostForm("id")

	_logTheCtx := fmt.Sprintf("%#v , %#v", ctx.Params, ctx.Keys)

	//检测，非管理员且id不匹配时
	if ok, err := CheckUserIsAdmin(ctx); !ok && !CheckUserIDIsCorrect(ctx, id) {
		if err != nil {
			logger.Logger.Error("获取用户信息失败", err, 0, _logTheCtx)
			ctx.Error(errors.GetError(errors.VALID_ERROR, "获取用户信息失败"))
			return
		}
		logger.Logger.Warn("用户非本人", 0, _logTheCtx)
		ctx.Error(errors.GetError(errors.VALID_ERROR, "用户非本人"))
		return
	}

	//删除
	if err := service.DeletePostService(id); err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse("ok", nil))

}

// UpdatePostHandle 更新评论
func UpdatePostHandle(ctx *gin.Context) {

	//获取参数
	data := &entity.Post{}
	_logTheCtx := fmt.Sprintf("%#v , %#v", ctx.Params, ctx.Keys)

	if err := ctx.ShouldBind(data); err != nil {
		logger.Logger.Error("获取update参数错误", err, 0, _logTheCtx)
		ctx.Error(errors.GetError(errors.INNER_ERROR, "update post 参数错误"))
		return
	}

	//检测,非管理员且id不匹配时
	if ok, err := CheckUserIsAdmin(ctx); !ok && !CheckUserIDIsCorrect(ctx, strconv.Itoa(data.UserId)) {
		if err != nil {
			logger.Logger.Error("获取用户信息失败", err, 0, _logTheCtx)
			ctx.Error(errors.GetError(errors.VALID_ERROR, "获取用户信息失败"))
			return
		}
		logger.Logger.Warn("用户非本人", 0, _logTheCtx)
		ctx.Error(errors.GetError(errors.VALID_ERROR, "用户非本人"))
		return
	}

	//更新
	if err := service.UpdatePostService(data); err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse("ok", nil))

}

// SelectPostHandle 查询post
func SelectPostHandle(ctx *gin.Context) {
	//获取参数
	id := ctx.Query("id")

	//无检测，任何有合法token的人都可以看到任意post

	//查询
	res := make(map[string]interface{})
	if err := service.SelectPostService(id, &res); err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse("ok", res))

}
