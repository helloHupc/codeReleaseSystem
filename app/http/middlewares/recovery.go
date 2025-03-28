package middlewares

import (
	"bytes"
	"codeReleaseSystem/pkg/logger"
	"context"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const timeFormat = "2006-01-02 15:04:05"

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 获取请求信息(不包含body)
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				timestamp := time.Now().Format(timeFormat)

				// 检查是否是连接中断
				isConnBroken := func() bool {
					if c.Request.Context().Err() == context.Canceled {
						return true
					}
					if ne, ok := err.(*net.OpError); ok {
						if se, ok := ne.Err.(*os.SyscallError); ok {
							errStr := strings.ToLower(se.Error())
							return strings.Contains(errStr, "broken pipe") ||
								strings.Contains(errStr, "connection reset by peer")
						}
					}
					return false
				}

				if isConnBroken() {
					// 连接中断日志
					logger.Logger.Warn("Client connection closed",
						zap.String("time", timestamp),
						zap.ByteString("request", httpRequest),
						zap.String("method", c.Request.Method),
						zap.String("path", c.Request.URL.Path),
						zap.String("ip", c.ClientIP()),
					)
					c.Abort()
					return
				}

				// 普通panic日志
				var body []byte
				if c.Request.Body != nil && c.Request.Method != http.MethodGet {
					body, _ = io.ReadAll(io.LimitReader(c.Request.Body, 1024))
					c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
				}

				logger.Logger.Error("Panic recovered",
					zap.String("time", timestamp),
					zap.Any("error", err),
					zap.ByteString("request", httpRequest),
					zap.ByteString("body", body),
					zap.Stack("stacktrace"),
				)

				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": "Internal Server Error",
				})
			}
		}()
		c.Next()
	}
}
