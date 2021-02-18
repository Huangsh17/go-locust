package main

import (
	"context"
	"encoding/json"
	"go-locust/contrib"
	"go-locust/dao"
	"go-locust/db"
	"go-locust/util"
	"go.uber.org/zap"
	"os"
	"time"
)

var (
	Key = "task"
)
var HostName string

var logger *zap.SugaredLogger

func init() {
	db.RedisInit()
	db.Connect()
	logger = util.DefaultLogger()
}

func main() {
	name, err := os.Hostname()
	if err != nil {
		logger.Errorw("get hostname fail", "error", err)
	}
	HostName = name
	go run()
	go register()
	go checkHeartbeat("")
	select {}
}

func getTask() {
	var locustTask dao.LocustTask
	conn := db.GetRedisConn()
	for {
		logger.Infow("开始监听任务队列")
		res, _ := conn.Do("BRPOP", Key, 0)
		value, _ := res.([]interface{})
		v, _ := value[1].([]byte)
		_ = json.Unmarshal(v, &locustTask)
		logger.Infow("测试任务", "task", locustTask)
		contrib.SendRequests(locustTask, context.Background())
	}
}

func register() {
	conn := db.GetRedisConn()
	_, err := conn.Do("SADD", "cluster", HostName)
	if err != nil {
		logger.Errorw("insert fail", "error", err)
	}
}

func checkHeartbeat(hostName string) {

	ticker := time.NewTicker(5 * time.Second)
	for {
		<-ticker.C
	}
}

func setCounter() {
	conn := db.GetRedisConn()
	_, _ = conn.Do("SET", HostName+"heartbeat")

}
func run() {
	getTask()
}
