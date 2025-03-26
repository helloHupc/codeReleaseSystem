package logger

import (
	"codeReleaseSystem/pkg/app"
	"codeReleaseSystem/pkg/config"
	"encoding/json"
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

func jsonString(value interface{}) string {
	b, err := json.Marshal(value)
	if err != nil {
		Logger.Error("json.Marshal failed", zap.Any("value", value), zap.Error(err))
	}
	return string(b)
}

// Dump 调试专用，不会中断程序，会在终端打印出 warning 消息。
// 第一个参数会使用 json.Marshal 进行渲染，第二个参数消息（可选）
func Dump(value interface{}, msg ...string) {
	valueString := jsonString(value)
	// 判断是否传第二个参数msg
	if len(msg) > 0 {
		Logger.Warn("Dump", zap.String(msg[0], valueString))
	} else {
		Logger.Warn("Dump", zap.String("value", valueString))
	}
}

// LogIf 当 err != nil 时记录 error 等级的日志
func LogIf(err error) {
	if err != nil {
		Logger.Error("Error Occurred:", zap.Error(err))
	}
}

// LogWarnIf 当 err != nil 时记录 warning 等级的日志
func LogWarnIf(err error) {
	if err != nil {
		Logger.Warn("Error Occurred:", zap.Error(err))
	}
}

// LogInfoIf 当 err != nil 时记录 info 等级的日志
func LogInfoIf(err error) {
	if err != nil {
		Logger.Info("Error Occurred:", zap.Error(err))
	}
}

// Debug 调试日志，详尽的程序日志
func Debug(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...)
}

// Info 告知类日志
func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

// Warn 警告类
func Warn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}

// Error 错误时记录，不应该中断程序，查看日志时重点关注
func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

// Fatal 级别同 Error(), 写完 log 后调用 os.Exit(1) 退出程序
func Fatal(msg string, fields ...zap.Field) {
	Logger.Fatal(msg, fields...)
}

// DebugString 记录一条字符串类型的 debug 日志
func DebugString(moduleName, name, msg string) {
	Logger.Debug(moduleName, zap.String(name, msg))
}

func InfoString(moduleName, name, msg string) {
	Logger.Info(moduleName, zap.String(name, msg))
}

func WarnString(moduleName, name, msg string) {
	Logger.Warn(moduleName, zap.String(name, msg))
}

func ErrorString(moduleName, name, msg string) {
	Logger.Error(moduleName, zap.String(name, msg))
}

func FatalString(moduleName, name, msg string) {
	Logger.Fatal(moduleName, zap.String(name, msg))
}

// DebugJSON 记录对象类型的 debug 日志，使用 json.Marshal 进行编码。调用示例：
//
//	logger.DebugJSON("Auth", "读取登录用户", auth.CurrentUser())
func DebugJSON(moduleName, name string, value interface{}) {
	Logger.Debug(moduleName, zap.String(name, jsonString(value)))
}

func InfoJSON(moduleName, name string, value interface{}) {
	Logger.Info(moduleName, zap.String(name, jsonString(value)))
}

func WarnJSON(moduleName, name string, value interface{}) {
	Logger.Warn(moduleName, zap.String(name, jsonString(value)))
}

func ErrorJSON(moduleName, name string, value interface{}) {
	Logger.Error(moduleName, zap.String(name, jsonString(value)))
}

func FatalJSON(moduleName, name string, value interface{}) {
	Logger.Fatal(moduleName, zap.String(name, jsonString(value)))
}
