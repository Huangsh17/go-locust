package main

import (
	"github.com/gin-gonic/gin"
	"go-locust/db"
	"go-locust/user"
)

func init() {
	db.RedisInit()
	db.Connect()
}
func main() {
	g := gin.Default()
	g.POST("/create_task", user.CreateTask)
	g.POST("/start_task", user.StartTask)
	g.POST("/stop_task", user.StopTask)
	_ = g.Run(":9999")
}
