package api

import (
	"github.com/XCPCBoard/api/errors"
	response "github.com/XCPCBoard/api/http"
	"github.com/XCPCBoard/user/entity"
	"github.com/XCPCBoard/user/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//UpdateWebAccountHandle 更新用户网站账户
func UpdateWebAccountHandle(ctx *gin.Context) {
	//获取参数
	data := &entity.Account{}
	if err := ctx.ShouldBind(data); err != nil {
		ctx.Error(errors.GetError(errors.INNER_ERROR, "参数错误"))
		return
	}

	//检查权限
	//若登录id和token id 不等
	if !CheckUserIDIsCorrect(ctx, strconv.Itoa(data.Id)) {
		//非超管
		if ok, err := CheckUserIsSuperAdmin(ctx); !ok {
			if err != nil {
				ctx.Error(errors.NewError(http.StatusInternalServerError, "验证管理员时错误:"+err.Error()))
				return
			}
			//若非管理员
			if ok, err := CheckUserIsAdmin(ctx); !ok {
				if err != nil {
					ctx.Error(errors.NewError(http.StatusInternalServerError, "验证管理员时错误:"+err.Error()))
					return
				}
				ctx.Error(errors.NewError(http.StatusForbidden, "用户无权限或token错误"))
				return
			} else if ok_, err_ := service.SelectUserIsAdmin(strconv.Itoa(data.Id)); ok_ { //被修改的用户是管理员
				if err_ != nil {
					ctx.Error(errors.NewError(http.StatusInternalServerError, "验证管理员时错误:"+err.Error()))
					return
				}
				ctx.Error(errors.NewError(http.StatusForbidden, "用户无权限修改其他管理员"))
			}
		}

	}

	//更新
	if err := service.UpdateWebAccountService(data); err != nil {
		ctx.Error(errors.GetError(errors.INNER_ERROR, "update web account error"))
		return
	} else {
		ctx.JSON(http.StatusOK, response.SuccessResponse("ok", nil))
	}
}

//SelectWebAccountHandle 查询用户网站账户
func SelectWebAccountHandle(ctx *gin.Context) {
	//获取id
	id := ctx.PostForm("id")

	//检测权限，若id不对
	if !CheckUserIDIsCorrect(ctx, id) {
		if ok, err := CheckUserIsAdmin(ctx); !ok { //非管理员
			if err != nil {
				ctx.Error(errors.NewError(http.StatusInternalServerError, "验证管理员时错误:"+err.Error()))
				return
			}
			ctx.Error(errors.NewError(http.StatusForbidden, "用户无权限或token错误"))
			return
		} //是管理员时就可以继续查
	}
	//更新
	user := map[string]interface{}{}
	err := service.SelectWebAccountService(id, &user)
	if err != nil {
		ctx.Error(errors.GetError(errors.INNER_ERROR, "select web account error"))
		return
	} else {
		ctx.JSON(http.StatusOK, response.SuccessResponse("ok", user))
	}
}
