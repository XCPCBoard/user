package api

import (
	"github.com/XCPCBoard/api/http/middleware"
	"github.com/gin-gonic/gin"
)

func BuildUserRouteEngine(engine *gin.Engine) {
	login := engine.Group("/login")
	{
		login.POST("/register", RegisterUser)
		login.POST("/user", UserLoginHandle)
		login.POST("/reset/password", ResetPassword)
		sent := login.Group("/sent")
		{
			em := sent.Group("/email")
			{
				//发送注册邮件
				em.POST("/register", SentRegisterEmailHandle)
				//发送重置密码邮件
				em.POST("/reset", SentResetPasswordEmailHandle)
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
			user.POST("/update", UpdateUserBySelfHandle)
			//管理员修改接口（超管可以修改一切）
			user.POST("/update/by/admin", UpdateUserByAdminHandle)
			user.POST("/delete", DeleteUserHandle)
			user.GET("/select", SelectUserHandle)
		}
		post := e.Group("/post")
		{
			post.POST("/create", CreatePostHandle)
			post.POST("/delete", DeletePostHandle)
			//管理员也可以修改
			post.POST("/update", UpdatePostHandle)
			post.GET("/select", SelectPostHandle)
		}
		account := e.Group("/webaccount")
		{
			account.POST("/update", UpdateWebAccountBySelfHandle)
			account.GET("/select", SelectWebAccountHandle)
		}
	}
}
