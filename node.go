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

var HostName string

var logger *zap.SugaredLogger

func init() {
	db.RedisInit()
	db.Connect()
	logger = util.DefaultLogger()
}

func main() {
	getHostName()
	go util.InitLog()
	go worker()
	go register()
	go ReportHeartbeatData()
	select {}
}

func worker() {
	var locustTask dao.LocustTask
	conn := db.GetRedisConn()
	for {
		logger.Infow("开始监听任务队列", "queue", HostName)
		res, _ := conn.Do("BRPOP", HostName+"_task", 0)
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

func ReportHeartbeatData() {
	conn := db.GetRedisConn()
	ticker := time.NewTicker(3 * time.Second)
	_, _ = conn.Do("SET", HostName+"_"+"heartbeat", 1)
	for {
		<-ticker.C
		_, _ = conn.Do("INCR", HostName+"_"+"heartbeat")
	}
}

func getHostName() {
	name, err := os.Hostname()
	if err != nil {
		logger.Errorw("get hostname fail", "error", err)
	}
	HostName = name
}
