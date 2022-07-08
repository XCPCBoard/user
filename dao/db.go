package dao

import (
	"XCPCer_board/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

var DBClient *sql.DB

const mysqlDriver = "mysql"

//NewDBClient 初始化db连接
func NewDBClient() (*sql.DB, error) {
	// 判断是否存在配置
	mysqlConfig := config.Conf.Storages[mysqlDriver]
	// 初始化连接
	dbClient, err := sql.Open(mysqlDriver, fmt.Sprintf("%v:%v@tcp(%v)/", mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host))
	if err != nil {
		log.Errorf("Open Sql Error: %v", err)
		return nil, err
	}
	return dbClient, nil
}
