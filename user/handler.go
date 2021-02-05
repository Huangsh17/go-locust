package user

import (
	"github.com/gin-gonic/gin"
	"go-locust/contrib"
	"go-locust/dao"
	"net/http"
	"runtime"
)

type Req struct {
	ThreadCount int    `form:"thread_count" json:"thread_count" binding:"required"`
	Method      string `form:"method" json:"method" binding:"required"`
	Url         string `form:"url" json:"url" binding:"required"`
	Body        string `form:"body" json:"body"`
	LoopCount   int    `form:"loop_count" json:"loop_count" binding:"required"`
}

func CreateTask(ctx *gin.Context) {
	var req Req
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"参数有误或者格式不对": err})
	}
	err := dao.CreateTask(req.ThreadCount, req.LoopCount, req.Method, req.Url, req.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err})
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func StartTask(ctx *gin.Context) {
	num := runtime.NumGoroutine()
	contrib.TaskQueue = make(chan dao.LocustTask)
	taskId, _ := ctx.GetPostForm("task_id")
	task := dao.QueryTask(taskId)
	go contrib.InitLocust()
	go func() {
		contrib.TaskQueue <- task
	}()
	ctx.JSON(http.StatusOK, gin.H{"msg": "start success", "g": num})
}

func StopTask(ctx *gin.Context) {
	_, _ = ctx.GetPostForm("task_id")
	contrib.StopTask()
	ctx.JSON(http.StatusOK, gin.H{"msg": "stop success"})
}

func TaskList(ctx *gin.Context) {

}

func TestApi(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "test success",
	})
}
