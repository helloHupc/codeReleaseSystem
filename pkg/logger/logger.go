package logger

import (
	"codeReleaseSystem/pkg/app"
	"codeReleaseSystem/pkg/config"
	"fmt"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger 全局日志对象 项目中所有地方都可以使用 Logger.xxx() 记录日志
var Logger *zap.Logger

func InitLogger() {
	// 获取日志配置
	logCfg := config.GetStringMapString("log")

	// 设置日志级别
	var level zapcore.Level
	if err := level.UnmarshalText([]byte(logCfg["level"])); err != nil {
		// 使用默认debug级别并记录错误
		level = zapcore.DebugLevel
		fmt.Printf("日志初始化错误，无效的日志级别 '%s', 请检查config/log.go配置\n", logCfg["level"])
	}

	// 初始化core
	core := zapcore.NewCore(
		getEncoder(),
		getLogWriter(
			logCfg["filename"],
			config.GetInt("log.max_size"),
			config.GetInt("log.max_backup"),
			config.GetInt("log.max_age"),
			config.GetBool("log.compress"),
			logCfg["type"],
		),
		level,
	)

	// 创建Logger
	Logger = zap.New(core,
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
}

// getEncoder 获取日志编码器
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// getLogWriter 获取日志写入器
func getLogWriter(filename string, maxSize int, maxBackups int, maxAge int, compress bool, logType string) zapcore.WriteSyncer {
	// 如果配置了按照日期记录日志文件
	if logType == "daily" {
		logname := time.Now().Format("2006-01-02") + ".log"
		filename = strings.ReplaceAll(filename, "logs.log", logname)
	}

	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   compress,
	}

	// 本地环境增加终端输出
	if app.IsLocal() {
		return zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(lumberJackLogger),
			zapcore.AddSync(os.Stdout),
		)
	}

	return zapcore.AddSync(lumberJackLogger)
}
