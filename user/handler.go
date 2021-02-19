package user

import (
	"github.com/gin-gonic/gin"
	"go-locust/contrib"
	"go-locust/dao"
	"net/http"
	"strconv"
)

type CreateTaskReq struct {
	ThreadCount int    `form:"thread_count" json:"thread_count" binding:"required"`
	Method      string `form:"method" json:"method" binding:"required"`
	Url         string `form:"url" json:"url" binding:"required"`
	Body        string `form:"body" json:"body"`
	LoopCount   int    `form:"loop_count" json:"loop_count" binding:"required"`
}

type StartTaskReq struct {
	TaskId               int `form:"task_id" json:"task_id" binding:"required"`
	OperatingEnvironment int `form:"operating_environment" json:"operating_environment" binding:"required"`
}

func CreateTask(ctx *gin.Context) {
	var req CreateTaskReq
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
	contrib.TaskQueue = make(chan dao.LocustTask)
	var req StartTaskReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"参数有误或者格式不对": err})
	}
	task := dao.QueryTask(strconv.Itoa(req.TaskId))
	switch req.OperatingEnvironment {
	case 1:
		go contrib.InitLocust()
		go func() {
			contrib.TaskQueue <- task
		}()
		ctx.JSON(http.StatusOK, gin.H{"msg": "start success"})
	case 2:
		isEmpty := dao.IsEmptyQueue()
		if isEmpty {
			ctx.JSON(http.StatusOK, gin.H{"msg": "任务池里面有任务了，等压测完在添加"})
			return
		}
		dao.AddTask(task)
		ctx.JSON(http.StatusOK, gin.H{"msg": "start success"})
	default:
	}
}

func StopTask(ctx *gin.Context) {
	_, _ = ctx.GetPostForm("task_id")
	ctx.JSON(http.StatusOK, gin.H{"msg": "stop success"})
}

func TaskList(ctx *gin.Context) {

}

func TestApi(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "test success",
	})
}
