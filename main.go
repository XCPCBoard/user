package main

import (
	"fmt"
	_ "github.com/FengZhg/go_tools/gin_logrus"
	"github.com/XCPCBoard/user/controller"
	"github.com/XCPCBoard/user/dao"
	_ "github.com/XCPCBoard/user/dao"
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
	a, b := controller.SelectPostController("1")
	fmt.Printf("%v,%v", a, b)
}

func init() {
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

}
