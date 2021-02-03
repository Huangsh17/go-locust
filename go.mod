module go-locust

go 1.13

require (
	//gitlab.dm-ai.cn/hexianmin/dm_go_micro_service v0.0.0-20210202111453-8aba785ad272
	github.com/gin-gonic/gin v1.6.3
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/satori/go.uuid v1.2.0
	gitlab.dm-ai.cn/zengzhiyong/dm-micro-service v0.0.0-20201109062705-07e82b4bb7bc // indirect
	gitlab.dm-ai.cn/zengzhiyong/dmai-apollo-client-go v0.0.0-20200511095548-1f98e7014942 // indirect
	go.uber.org/zap v1.16.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

//replace gitlab.dm-ai.cn/hexianmin/dm_go_micro_service v0.0.0 => dm_go_micro_service v0.0.0
