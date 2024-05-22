package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// NewLoggerConfig returns a gin.Engine with custom logger configuration
func NewLoggerConfig(logFilePath string, skipPaths []string) (*gin.Engine, error) {
	// 打开日志文件
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("无法打开日志文件: %v", err)
	}

	// 配置 gin 的 LoggerConfig
	loggerConfig := gin.LoggerConfig{
		Formatter: func(param gin.LogFormatterParams) string {
			return fmt.Sprintf("[GIN] %v | %3d | %13v | %15s | %-7s %s\n%s",
				param.TimeStamp.Format(time.RFC1123),
				param.StatusCode,
				param.Latency,
				param.ClientIP,
				param.Method,
				param.Path,
				param.ErrorMessage,
			)
		},
		Output:    logFile, // 将日志输出到文件
		SkipPaths: skipPaths,
	}

	// 创建 gin 服务器并使用自定义日志配置
	r := gin.New()
	r.Use(gin.LoggerWithConfig(loggerConfig))

	return r, nil
}
