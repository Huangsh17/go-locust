package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Req struct {
	ThreadCount int    `form:"thread_count" json:"thread_count" binding:"required"`
	Method      string `form:"method" json:"method" binding:"required"`
	Url         string `form:"url" json:"url" binding:"required"`
	Body        string `form:"body" json:"body" binding:"required"`
}

func CreateTask(ctx *gin.Context) {
	var req Req

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"参数有误或者格式不对": err})
	}

}

func StartTask(ctx *gin.Context) {

}

func StopTask(ctx *gin.Context) {

}
