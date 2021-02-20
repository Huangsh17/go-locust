package dao

import (
	"encoding/json"
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

func CreateTask(threadCount, loopCount int, method, url, body string) error {
	sql := "insert into locust_task (thread_count,method,url,body,loop_count)values (?,?,?,?,?)"
	_, err := db.Conn.Exec(sql, threadCount, method, url, body, loopCount)
	if err != nil {
		util.Sugar.Errorw("insert fail", "error", err)
	}
	return nil
}

func QueryTask(taskId string) (locustTask LocustTask) {
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

func AddTask(lt LocustTask) {
	conn := db.GetRedisConn()
	res, _ := json.Marshal(lt)
	nodes := GetHostName()
	for _, node := range nodes {

		_, err := conn.Do("LPUSH", node+"_"+"task", res)
		if err != nil {
			util.Sugar.Errorw("lpush fail", "error", err)
		}
	}
}

func GetHostName() []string {
	conn := db.GetRedisConn()
	nodes := make([]string, 0)
	reply, _ := conn.Do("SMEMBERS", "cluster")
	values, _ := reply.([]interface{})
	for _, value := range values {
		nodes = append(nodes, string(value.([]byte)))
	}
	return nodes
}

func IsEmptyQueue() bool {
	conn := db.GetRedisConn()
	nodes := GetHostName()
	j := 0
	for _, node := range nodes {
		length, err := conn.Do("LLEN", node+"_task")
		if err != nil {
			util.Sugar.Errorw("LLEN fail", "error", err)
		}
		if length.(int64) >= 1 {
			j++
		}
	}
	if 0 < j && j <= len(nodes) {
		return true
	}
	return false
}
