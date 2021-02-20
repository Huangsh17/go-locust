package config

// 项目日志配置
type LogBasicsConfig struct {
	// 服务名称
	ServerName string
	// 等级
	Level string
	// 输出文件路径 默认在当前项目根目录生成logs目录下存放日志
	OutPath string
	// 编码格式，支持console，json
	Encoding string
	// 输出方式 0:输出控制台 1:输出文件 2 ：同时输出文件和控制台;文件输出路径不用填写
	OutputType int
	// 使用配置中心 以上所有字段都不用填写
	//ConfigCore bool
	//// 配置中心链接
	//ConfigCoreUrl string
}

// 日志示例配置例子
var Logging = []LogBasicsConfig{
	// 铭感词过滤功能模块日志打印
	{"sensitive", "info", "sensitive.log", "json", 1},
}

// db配置
//mysql
const (
	HOST_MYSQL     = "127.0.0.1:3306"
	USER_MYSQL     = "root"
	PASSWORD_MYSQL = "123"
	DB_NAME_MYSQL  = "hsh"
)

// redis
const (
	HOST_REDIS     = "127.0.0.1:6379"
	USER_REDIS     = ""
	PASSWORD_REDIS = ""
)

// etcd
const (
	HOST_ETCD = "127.0.0.1:6379"
	USER_ETCD = ""
)
