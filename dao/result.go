package dao

import (
	"go-locust/db"
	"go-locust/util"
	"time"
)

type LocustResult struct {
	ID         int       `json:"id"`
	TaskId     int       `json:"task_id"`
	Body       int       `json:"body"`
	CreateTime time.Time `json:"create_time"`
}

func CreateResult(string2 string, taskId int) error {
	sql := "insert into locust_result (body,task_id,create_time)values (?,?,?)"
	_, err := db.Conn.Exec(sql, string2, taskId, time.Now())
	if err != nil {
		util.Sugar.Errorw("insert fail", "error", err)
	}
	return nil
}
