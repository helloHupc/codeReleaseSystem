package logger

import (
	"testing"
)

func TestLogger(t *testing.T) {
	// 初始化日志
	InitLogger()

	// 记录测试日志
	Logger.Info("This is an info message")
	Logger.Debug("This is a debug message")
	Logger.Warn("This is a warning message")
	Logger.Error("This is an error message")
}
