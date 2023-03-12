package handle

import (
	"fmt"
	response "github.com/XCPCBoard/api/http"
	"github.com/XCPCBoard/api/http/middleware"
	"github.com/XCPCBoard/common/errors"
	"github.com/XCPCBoard/common/logger"
	"github.com/XCPCBoard/user/entity"
	"github.com/XCPCBoard/user/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

// DeletePostByAdminHandle 管理员删除评论
func DeletePostByAdminHandle(ctx *gin.Context) {
	//获取参数
	id := ctx.PostForm("id")

	_logTheCtx := fmt.Sprintf("%#v , %#v", ctx.Params, ctx.Keys)

	//删除
	if err := service.DeletePostService(id); err != nil {
		logger.L.Error("删除评论错误", err, 0, _logTheCtx)
		ctx.Error(err)
		return
	}

	response.SuccessResponseAddToken(ctx, nil)

}

// DeleteUserHandle 超管删除用户
func DeleteUserHandle(ctx *gin.Context) {

	id := ctx.PostForm("id")
	_logTheCtx := fmt.Sprintf("%#v , %#v", ctx.Params, ctx.Keys)

	uid, err := strconv.Atoi(id)
	if err != nil {
		logger.L.Warn("用户id不为int"+err.Error(), 0, id)
		ctx.Error(errors.GetError(errors.VALID_ERROR, "id 错误"))
		return
	}
	if uid, ok := ctx.Get(middleware.TokenIDStr); !ok {
		e := errors.CreateError(errors.INNER_ERROR.Code, "获取用户token中的id出错", _logTheCtx)
		logger.L.Err(e, 0)
		ctx.Error(e)
		return
	} else if uid == id {
		e := errors.CreateError(errors.ERROR.Code, "超管不能删除自己", _logTheCtx)
		logger.L.Err(e, 0)
		ctx.Error(e)
		return
	}

	//删除
	user := entity.Users{
		Id: uid,
	}
	if err := service.DeleteUserService(&user); err != nil {
		ctx.Error(err)
		return
	}

	response.SuccessResponseAddToken(ctx, nil)

}

// UpdateUserByAdminHandle 更新用户数据,by 超管
func UpdateUserByAdminHandle(ctx *gin.Context) {

	uid := ctx.PostForm("id")

	user := entity.Users{
		Name:        ctx.PostForm("name"),
		ImagePath:   ctx.PostForm("image"),
		Signature:   ctx.PostForm("signature"),
		PhoneNumber: ctx.PostForm("phone"),
		QQNumber:    ctx.PostForm("qq"),
	}

	if err := service.UpdateUserService(uid, &user); err != nil {
		ctx.Error(err)
		return
	}
	response.SuccessResponseAddToken(ctx, nil)
}

// UpdateWebAccountByAdminHandle 管理员更新用户爬虫网站账户信息
func UpdateWebAccountByAdminHandle(ctx *gin.Context) {
	_logTheCtx := fmt.Sprintf("%#v , %#v", ctx.Params, ctx.Keys)
	//获取参数
	data := &entity.Accounts{}
	if err := ctx.ShouldBind(data); err != nil {
		logger.L.Error("获取更新网站账户时，参数错误", err, 0, _logTheCtx)
		ctx.Error(errors.GetError(errors.VALID_ERROR, "参数错误"))
		return
	}

	//更新
	if err := service.UpdateWebAccountWithOutNumService(data); err != nil {
		ctx.Error(err)
		return
	}

	response.SuccessResponseAddToken(ctx, nil)

}
