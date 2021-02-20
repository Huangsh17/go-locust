package contrib

import (
	"context"
	"go-locust/dao"
	"go-locust/util"
	"sync"
)

var (
	http *util.HttpClient
	_    context.CancelFunc
	wg   sync.WaitGroup
	// 本地单机任务队列
	TaskQueue chan dao.LocustTask
)

func SendRequests(task dao.LocustTask, ctx context.Context) {
	startTask(task, ctx)
}

func startTask(task dao.LocustTask, ctx context.Context) {
	for i := 1; i <= task.LoopCount; i++ {
		for j := 1; j <= task.ThreadCount; j++ {
			wg.Add(1)
			go locust(task)
		}
	}
	wg.Wait()
}

func locust(task dao.LocustTask) string {
	switch task.Method {
	case "get":
		resp, _ := http.Get(task.Url)
		util.Sugar.Infow("执行成功", "result", resp)
		_ = dao.CreateResult(resp, task.ID)
		wg.Done()
		return resp
	case "post":
		resp, _ := http.Post(task.Url, task.Body)
		util.Sugar.Infow("执行成功", "result")
		_ = dao.CreateResult(resp, task.ID)
		wg.Done()
		return resp
	}
	return ""
}

func InitLocust() {
	ctx, cancel := context.WithCancel(context.Background())
	_ = cancel
	for {
		select {
		case <-ctx.Done():
			util.Sugar.Infow("goroutine quit")
		case task := <-TaskQueue:
			SendRequests(task, ctx)
			return
		}
	}
}
