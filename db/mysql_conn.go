package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-locust/config"
	"go-locust/util"
	"time"
)

var (
	Conn *sql.DB
)

func Connect() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Local", config.USER_MYSQL, config.PASSWORD_MYSQL, config.HOST_MYSQL, config.DB_NAME_MYSQL)
	conn, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		util.Sugar.Errorw("mysql connection fail", "error", err)
	}
	// 设置连接的最大生存时间，以确保连接可以被驱动安全关闭。官方建议小于5分钟。
	conn.SetConnMaxLifetime(time.Minute * 3)
	// 设置打开的最大连接数，取决于mysql服务器和具体应用程序
	conn.SetMaxOpenConns(10)
	// 设置最大闲置连接数，这个连接数应大于等于打开的最大连接，否则需要额外连接时会频繁进行打开关闭。
	// 最好与最大连接数保持相同，当大于最大连接数时，内部自动会减少到与最大连接数相同。
	conn.SetMaxIdleConns(10)
	Conn = conn
}
