package contrib

import (
	"context"
	"fmt"
	"go-locust/dao"
	"go-locust/util"
	"sync"
)

var (
	http   *util.HttpClient
	Cancel context.CancelFunc
	wg     sync.WaitGroup
)

func SendRequests(task dao.LocustTask) {
	ctx, cancel := context.WithCancel(context.Background())
	startTask(task, ctx)
	Cancel = cancel
}

func startTask(task dao.LocustTask, ctx context.Context) {

	for i := 1; i <= task.LoopCount; i++ {
		for j := 0; j <= task.ThreadCount; j++ {
			go locust(task, ctx)
		}
	}
}

func locust(task dao.LocustTask, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			util.Sugar.Infow("goroutine quit")
			return
		default:
			switch task.Method {
			case "get":
				resp, _ := http.Get(task.Url)
				fmt.Println(resp)
				return
			case "post":
				_, _ = http.Post(task.Url, task.Body)
			}

		}
	}
}

func StopTask() {
	Cancel()
}
