package api

import (
	"github.com/XCPCBoard/api/http/middleware"
	"github.com/XCPCBoard/user/api/handle"
	"github.com/XCPCBoard/user/api/mid"
	"github.com/XCPCBoard/user/service"
	"github.com/gin-gonic/gin"
)

func BuildUserRouteEngine(engine *gin.Engine) {
	login := engine.Group("/login")
	{
		login.POST("/register", handle.RegisterUser)
		login.POST("/user", handle.UserLoginHandle)
		login.POST("/reset/password", handle.ResetPassword)
		sent := login.Group("/sent")
		{
			em := sent.Group("/email")
			{
				//发送注册邮件
				em.POST("/register", service.SentRegisterEmailHandle)
				//发送重置密码邮件
				em.POST("/reset", service.SentResetPasswordEmailHandle)
			}
			//sent.GET("/ver")

		}
	}

	//加上token验证
	e := engine.Group("/api",
		middleware.AuthMiddleware())
	{
		user := e.Group("/user")
		{
			//更新自己的数据
			user.POST("/u", handle.UpdateUserBySelfHandle)

			s := user.Group("/s")
			{
				s.GET("/id", handle.SelectUserHandle)
				//模糊查询
				s.GET("/name", handle.SelectUserByNameHandle)
				s.GET("/account", handle.SelectUserByAccountHandle)
			}

		}
		post := e.Group("/post")
		{
			post.POST("/c", handle.CreatePostHandle)
			post.POST("/d", handle.DeleteSelfPostHandle)
			//改自己的
			post.POST("/u", handle.UpdateSelfPostHandle)
			post.GET("/s", handle.SelectPostHandle)

		}
		account := e.Group("/webaccount")
		{
			account.POST("/u", handle.UpdateWebAccountBySelfHandle)
			account.GET("/s", handle.SelectWebAccountHandle)
			//分页查询
			account.GET("/all", handle.SelectAllUserAccountDataPaging)
		}

		//超管权限
		super := e.Group("/super", mid.SuperAdminAuthMiddleware())
		{
			super.POST("/post/d", handle.DeletePostByAdminHandle)
			super.POST("/user/d", handle.DeleteUserHandle)
			super.POST("/user/u", handle.UpdateUserByAdminHandle)
			super.POST("/webaccount/u", handle.UpdateWebAccountByAdminHandle)
		}
	}
}
