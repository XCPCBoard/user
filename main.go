package main

import (
	"fmt"
	_ "github.com/FengZhg/go_tools/gin_logrus"
	"github.com/XCPCBoard/api/http/middleware"
	"github.com/XCPCBoard/common/config"
	"github.com/XCPCBoard/common/dao"
	_ "github.com/XCPCBoard/common/dao"
	"github.com/XCPCBoard/common/errors"
	"github.com/XCPCBoard/common/logger"
	"github.com/XCPCBoard/common/mail"
	"github.com/XCPCBoard/user/api"
	"github.com/XCPCBoard/user/entity"
	"github.com/XCPCBoard/user/service"
	"github.com/gin-gonic/gin"
)

// 主入口函数
func main() {

	//a, b := controller.CreatUserController("test2", "123456", "12")
	//a, b := controller.DeleteUserController("121", "test1")
	//a, b := controller.UpdateUserController(12, map[string]interface{}{
	//	"name": "testA",
	//})
	//a, b := controller.UpdateUserController(map[string]interface{}{
	//	"id":   "13",
	//	"name": "6666",
	//})
	//a, b := api.SelectPostController("1")
	//fmt.Printf("%v,%v", a, b)

	if err := service.HandleAllAccount(2, func(user *entity.Account) *errors.MyError {
		fmt.Println(fmt.Sprintf("%#v", user))
		return nil
	}); err != nil {
		panic(err)
	}

}

func initEngine() {
	engine := gin.Default()

	engine.Use(middleware.LoggerToFile())
	engine.Use(middleware.ErrorHandler())

	api.BuildUserRouteEngine(engine)
	engine.Run()

}

func init() {
	//依赖包

	config.BuildConfig("./config.yaml")

	//dao
	redisClient, err := dao.NewRedisClient()
	if err != nil {
		panic(err)
	}
	dao.RedisClient = redisClient
	dbClient, err := dao.NewDBClient()
	if err != nil {
		panic(err)
	}
	dao.DBClient = dbClient

	//log
	if err = logger.InitLogger(); err != nil {
		panic(err)
	}
	//email
	mail.InitEmail()

}
