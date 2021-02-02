package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-locust/config"
	"go-locust/util"
)

var (
	Conn *sql.DB
)

func Connect() {
	addr := fmt.Sprintf("%s:%s@tcp(%s)/%s", config.USER_MYSQL, config.PASSWORD_MYSQL, config.HOST_MYSQL, config.DB_NAME_MYSQL)
	conn, err := sql.Open("mysql", addr)
	if err != nil {
		util.Sugar.Errorw("mysql connection fail", "error", err)
	}
	Conn = conn
}
