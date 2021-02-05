package contrib

import (
	"context"
	"fmt"
	"go-locust/dao"
	"go-locust/util"
	"sync"
)

var (
	http      *util.HttpClient
	Cancel    context.CancelFunc
	wg        sync.WaitGroup
	TaskQueue chan dao.LocustTask
)

func sendRequests(task dao.LocustTask, ctx context.Context) {
	startTask(task, ctx)
}

func startTask(task dao.LocustTask, ctx context.Context) {
	for i := 1; i <= task.LoopCount; i++ {
		for j := 0; j <= task.ThreadCount; j++ {
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
		fmt.Println(resp)
		_ = dao.CreateResult(resp, task.ID)
		wg.Done()
		return resp
	case "post":
		resp, _ := http.Post(task.Url, task.Body)
		_ = dao.CreateResult(resp, task.ID)
		wg.Done()
		return resp
	}
	return ""
}

func InitLocust() {
	ctx, cancel := context.WithCancel(context.Background())
	Cancel = cancel
	for {
		select {
		case <-ctx.Done():
			util.Sugar.Infow("goroutine quit")
		case task := <-TaskQueue:
			sendRequests(task, ctx)
			return
		}
	}
}

func StopTask() {
	Cancel()
}
