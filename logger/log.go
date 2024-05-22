package logger

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func init() {
	// 设置日志输出到文件
	file, err := os.OpenFile("logger/access.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Fatalf("无法打开日志文件: %v", err)
	}
	logger.SetOutput(file)

	// 设置日志格式
	logger.SetFormatter(&logrus.JSONFormatter{})

	// 设置日志级别
	logger.SetLevel(logrus.InfoLevel)
}

// GinLogMiddleware 日志中间件
func GinLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 处理请求
		c.Next()

		// 记录请求日志
		logger.WithFields(logrus.Fields{
			"timestamp": start.Format(time.RFC3339),
			"method":    c.Request.Method,
			"url":       c.Request.URL.Path,
			"client_ip": c.ClientIP(),
			"status":    c.Writer.Status(),
			"duration":  time.Since(start),
		}).Info("访问日志")
	}
}
