// 该模块封装了zap框架，提供日志层面的基本配置，包括堆栈，日志切割，备份，多实例化日志
// 使用方法：直接调用NewLogger方法，生成logger对象，传入serverName(config配置的serverName（必须和serverName名字一样，不然会生成对象为nil）)
// 使用goland的grep-console插件输出不同级别颜色的日志
package util

import (
	"context"
	"errors"
	"go-locust/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var (
	// 多实例服务配置,配置日志基础配置，比如等级，服务名称，保存路径等等
	LogConfigs []config.LogBasicsConfig
	// 日志内容基础配置
	EncoderConfig zapcore.EncoderConfig
)

var Sugar *zap.SugaredLogger

func init() {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		FunctionKey:    "function",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
	}
	LogConfigs = config.Logging
	EncoderConfig = encoderConfig
}

func InitLogger(logConfig config.LogBasicsConfig) *zap.SugaredLogger {
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	switch logConfig.Level {
	case "info":
		atomicLevel.SetLevel(zap.InfoLevel)
	case "debug":
		atomicLevel.SetLevel(zap.DebugLevel)
	case "error":
		atomicLevel.SetLevel(zap.ErrorLevel)
	}
	// 初始化日志回滚，切割，备份配置
	var hook lumberjack.Logger
	if len(logConfig.OutPath) == 0 && logConfig.OutputType == 0 {
		hook = lumberjack.Logger{}
	} else {
		hook = lumberjack.Logger{
			Filename:   "./logs/" + logConfig.OutPath, // 日志文件路径
			MaxSize:    200,                           // 每个日志文件保存的最大尺寸 单位：M
			MaxBackups: 30,                            // 日志文件最多保存多少个备份
			//MaxAge:     7,                        // 文件最多保存多少天
			Compress: false, // 是否压缩
		}
	}
	var outType zapcore.WriteSyncer
	switch logConfig.OutputType {
	case 0:
		outType = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)) // 输出控制台
	case 1:
		outType = zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook)) // 输出文件
	case 2:
		outType = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)) // 同时输出控制台和文件
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(EncoderConfig), // 编码器配置
		outType,
		atomicLevel, // 日志级别
	)
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	filed := zap.Fields(zap.String("serviceName", logConfig.ServerName))
	// 构造日志对象
	logger := zap.New(core, caller, development, filed, zap.AddStacktrace(zap.WarnLevel))
	defer logger.Sync()
	sugar := logger.Sugar()
	return sugar
}

func NewLogger(serverName string) (*zap.SugaredLogger, error) {
	for _, logConfig := range LogConfigs {
		if logConfig.ServerName == serverName {
			return InitLogger(logConfig), nil
		}
	}
	return nil, errors.New("serverName和你配置的不一样，请用config.go里面配置的serverName")
}

func DefaultLogger() *zap.SugaredLogger {
	logConfig := config.LogBasicsConfig{
		ServerName: "default",
		Level:      "info",
		OutPath:    "default.log",
		Encoding:   "json",
		OutputType: 2,
	}
	return InitLogger(logConfig)
}

func InitLog() {
	logger := DefaultLogger()
	Sugar = logger
}

// 日志业务流追踪，在需要追踪一个完整的业务流开始前调用，并并写道日志里面，比如：
//ctx := TraceLogger()
//util.Sugar.Infow("购物流程","traceId",ctx.Value("traceId"))
func TraceLogger() context.Context {
	ctx := context.WithValue(context.Background(), "traceId", GenUUID())
	return ctx
}
