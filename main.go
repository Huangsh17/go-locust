package main

import (
	"github.com/gin-gonic/gin"
	"go-locust/db"
	"go-locust/user"
	"go-locust/util"

	"github.com/gin-contrib/pprof"
)

func init() {
	initLogger()
	db.RedisInit()
	db.Connect()
}
func main() {
	g := gin.Default()
	pprof.Register(g)
	g.POST("/create_task", user.CreateTask)
	g.POST("/start_task", user.StartTask)
	g.POST("/stop_task", user.StopTask)
	g.GET("/task_list", user.TaskList)
	g.GET("/test", user.TestApi)
	_ = g.Run(":9999")
}

func initLogger() {
	logger := util.DefaultLogger()
	util.Sugar = logger
}
