package main

import (
	"context"
	"encoding/json"
	"go-locust/contrib"
	"go-locust/dao"
	"go-locust/db"
	"go-locust/util"
	"go.uber.org/zap"
)

var (
	Key = "task"
)

var logger *zap.SugaredLogger

func init() {
	db.RedisInit()
	db.Connect()
	logger = util.DefaultLogger()
}

func main() {
	go run()
	go clusterStatus()
	select {}
}

func getTask() {
	var locustTask dao.LocustTask
	conn := db.GetRedisConn()
	for {
		res, _ := conn.Do("BRPOP", Key, 0)
		value, _ := res.([]interface{})
		v, _ := value[1].([]byte)
		_ = json.Unmarshal(v, &locustTask)
		logger.Infow("测试任务", "task", locustTask)
		contrib.SendRequests(locustTask, context.Background())
	}
}

func clusterStatus() {
	//name, err := os.Hostname()
	//if err != nil{
	//	logger.Errorw("get hostname fail","error",err)
	//}
	//conn := db.GetRedisConn()
	//conn.Do("")

}

func run() {
	getTask()
}
