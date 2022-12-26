package api

import (
	"github.com/XCPCBoard/user/entity"
	"github.com/XCPCBoard/user/service"
	response "github.com/XCPCBoard/utils/http"
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

	uid, err := strconv.Atoi(userId)
	if err != nil {
		log.Errorf("userId is not int type :%v", err.Error())
		ctx.JSON(http.StatusForbidden, response.FailResponse("userId is not int type", nil))
	}
	post := entity.Post{
		UserId:  uid,
		Title:   title,
		Content: content,
		Note:    note,
	}
	if err := service.CreatePostService(&post); err != nil {
		ctx.JSON(http.StatusForbidden, response.FailResponse("creat post error", nil))
	} else {
		ctx.JSON(http.StatusOK, response.SuccessResponse("ok", nil))
	}
}

//DeletePostHandle 删除评论
func DeletePostHandle(ctx *gin.Context) {
	id := ctx.PostForm("id")

	if err := service.DeletePostService(id); err != nil {
		ctx.JSON(http.StatusForbidden, response.FailResponse("delete post error", nil))
	} else {
		ctx.JSON(http.StatusOK, response.SuccessResponse("ok", nil))
	}
}

//UpdatePostHandle 更新评论
func UpdatePostHandle(ctx *gin.Context) {
	post := ctx.PostFormMap("post")

	if err := service.UpdatePostService(post); err != nil {
		ctx.JSON(http.StatusForbidden, response.FailResponse("update post error", nil))
	} else {
		ctx.JSON(http.StatusOK, response.SuccessResponse("ok", nil))
	}
}

//SelectPostHandle 查询post
func SelectPostHandle(ctx *gin.Context) {
	id := ctx.PostForm("id")

	res := map[string]interface{}{}
	if err := service.SelectPostService(id, &res); err != nil {
		ctx.JSON(http.StatusForbidden, response.FailResponse("update post error", nil))
	} else {
		ctx.JSON(http.StatusOK, response.SuccessResponse("ok", res))
	}
}
