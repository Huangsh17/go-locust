package dao

import (
	"go-locust/db"
	"go-locust/util"
)

type LocustTask struct {
	ID          int    `json:"id"`
	ThreadCount int    `json:"thread_count"`
	Method      string `json:"method"`
	Url         string `json:"url"`
	Body        string `json:"body"`
	LoopCount   int    `json:"loop_count"`
}

func Create(threadCount, loopCount int, method, url, body string) error {
	sql := "insert into locust_task (thread_count,method,url,body,loop_count)values (?,?,?,?,?)"
	_, err := db.Conn.Exec(sql, threadCount, method, url, body, loopCount)
	if err != nil {
		util.Sugar.Errorw("insert fail", "error", err)
	}
	return nil
}

func Query(taskId string) (locustTask LocustTask) {
	sql := "select * from locust_task where id = ?"
	rows, err := db.Conn.Query(sql, taskId)
	if err != nil {
		util.Sugar.Errorw("select fail", "error", err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&locustTask.ID, &locustTask.ThreadCount, &locustTask.Method, &locustTask.Url, &locustTask.Body, &locustTask.LoopCount)
		if err != nil {
			util.Sugar.Errorw("scan fail", "error", err)
		}
	}
	return
}
