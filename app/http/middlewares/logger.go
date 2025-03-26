package middlewares

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"codeReleaseSystem/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type responseWriterWrapper struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *responseWriterWrapper) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 基础信息记录
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method
		isGet := method == http.MethodGet

		// 非GET请求读取请求体
		var requestBody []byte
		if !isGet && c.Request.Body != nil {
			requestBody, _ = io.ReadAll(io.LimitReader(c.Request.Body, 1024))
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		// 包装ResponseWriter
		writer := &responseWriterWrapper{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		if !isGet {
			c.Writer = writer
		}

		// 处理请求
		c.Next()

		// 准备日志字段
		latency := time.Since(start)
		status := c.Writer.Status()
		logFields := []zap.Field{
			zap.Int("status", status),
			zap.String("method", method),
			zap.String("path", path),
			zap.String("ip", c.ClientIP()),
			zap.Duration("latency", latency),
		}

		// 非GET请求添加body
		if !isGet {
			responseBody := ""
			if writer.body.Len() > 0 {
				responseBody = writer.body.String()
				if len(responseBody) > 1024 {
					responseBody = responseBody[:1024] + "...[truncated]"
				}
			}
			logFields = append(logFields,
				zap.ByteString("request_body", requestBody),
				zap.String("response_body", responseBody),
			)
		}

		// 按状态码记录日志
		switch {
		case status >= 500:
			logger.Logger.Error("Server error", logFields...)
		case status >= 400:
			logger.Logger.Warn("Client error", logFields...)
		default:
			logger.Logger.Info("Request processed", logFields...)
		}
	}
}
