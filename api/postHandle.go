package api

import (
	"github.com/XCPCBoard/api/errors"
	response "github.com/XCPCBoard/api/http"
	"github.com/XCPCBoard/user/entity"
	"github.com/XCPCBoard/user/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

//CreatePostHandle 创建评论
func CreatePostHandle(ctx *gin.Context) {

	userId := ctx.PostForm("userId")
	title := ctx.PostForm("title")
	content := ctx.PostForm("content")
	note := ctx.PostForm("note")

	//对比tokenID和userID
	if !CheckUserIDIsCorrect(ctx, userId) {
		ctx.Error(errors.NewError(http.StatusForbidden, "用户非本人限"))
		return
	}

	uid, err := strconv.Atoi(userId)
	if err != nil {
		log.Errorf("userId is not int type :%v", err.Error())
		ctx.Error(errors.GetError(errors.ILLEGAL_DATA, "userId is not int type "))
		return
	}
	post := entity.Post{
		UserId:  uid,
		Title:   title,
		Content: content,
		Note:    note,
	}
	if err := service.CreatePostService(&post); err != nil {
		ctx.Error(errors.GetError(errors.INNER_ERROR, "Create post error"))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse("ok", nil))
}

//DeletePostHandle 删除评论
func DeletePostHandle(ctx *gin.Context) {
	//获取参数
	id := ctx.PostForm("id")

	//检测，非管理员且id不匹配时
	if ok, err := CheckUserIsAdmin(ctx); !ok && !CheckUserIDIsCorrect(ctx, id) {
		if err != nil {
			ctx.Error(errors.NewError(http.StatusInternalServerError, "获取用户信息失败"))
			return
		}
		ctx.Error(errors.NewError(http.StatusForbidden, "用户非本人"))
		return
	}

	//删除
	if err := service.DeletePostService(id); err != nil {
		ctx.Error(errors.GetError(errors.INNER_ERROR, "delete post error"))
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse("ok", nil))

}

//UpdatePostHandle 更新评论
func UpdatePostHandle(ctx *gin.Context) {

	//获取参数
	data := &entity.Post{}
	if err := ctx.ShouldBind(data); err != nil {
		ctx.Error(errors.GetError(errors.INNER_ERROR, "update post 参数错误"))
		return
	}

	//检测,非超级管理员且id不匹配时
	if ok, err := CheckUserIsSuperAdmin(ctx); ok && !CheckUserIDIsCorrect(ctx, strconv.Itoa(data.UserId)) {
		if err != nil {
			ctx.Error(errors.NewError(http.StatusInternalServerError, "检测用户是否为管理员时错误"))
			return
		}
		ctx.Error(errors.NewError(http.StatusForbidden, "用户非本人限"))
		return
	}

	//更新
	if err := service.UpdatePostService(data); err != nil {
		ctx.Error(errors.GetError(errors.INNER_ERROR, "update post error"))
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse("ok", nil))

}

//SelectPostHandle 查询post
func SelectPostHandle(ctx *gin.Context) {
	//获取参数
	id := ctx.Query("id")

	//无检测，任何有合法token的人都可以看到任意post

	//查询
	res := map[string]interface{}{}
	if err := service.SelectPostService(id, &res); err != nil {
		ctx.Error(errors.GetError(errors.INNER_ERROR, "Select post error"))
		return
	} else {
		ctx.JSON(http.StatusOK, response.SuccessResponse("ok", res))
	}
}
