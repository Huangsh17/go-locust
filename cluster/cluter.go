package cluster

import (
	"go-locust/db"
	"go-locust/util"
	"strconv"
	"strings"
	"time"
)

func HealthCheck() {
	conn := db.GetRedisConn()
	keys, _ := conn.Do("KEYS", "*")
	values, _ := keys.([]interface{})
	for _, value := range values {
		v, _ := value.([]byte)
		if strings.Contains(string(v), "heartbeat") {
			go check(string(v))
		}
	}
}

func check(key string) {
	var initValue int = 1
	conn := db.GetRedisConn()
	ticker := time.NewTicker(5 * time.Second)
	genSplit := strings.Split(key, "_")
	for {
		<-ticker.C
		countStr, _ := conn.Do("GET", key)
		count, _ := strconv.Atoi(string(countStr.([]byte)))
		if initValue == count {
			_, err := conn.Do("SREM", "cluster", genSplit[0])
			if err != nil {
				util.Sugar.Errorw("SREM fail", "error", err)
			}
			util.Sugar.Infow("集群健康状态", "status", "不健康,已从redis剔除该节点", "node", genSplit[0])
			continue
		}
		initValue = count
		util.Sugar.Infow("集群健康状态", "status", "健康", "node", genSplit[0])
	}
}

// 集群控制器
func LocustController() {

}
