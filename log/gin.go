package log

import (
	"github.com/clearcodecn/dep/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

var (
	logger *logrus.Logger
)

func init() {
	logger = logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "01-02 15:04:05",
	})
	logger.SetReportCaller(true)
}

type GinLoggerConfig struct {
	SkipPrefixes []string
}

func GinLogger(config GinLoggerConfig) gin.HandlerFunc {
	var (
		skipPrefixes []string
	)
	skipPrefixes = append(config.SkipPrefixes, "/favicon.ico")
	skipPrefixes = append(config.SkipPrefixes, "/static")

	// 生产环境下，打印req body 但是不打印 res body
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()
		// 处理请求
		c.Next()

		for _, v := range skipPrefixes {
			if strings.HasPrefix(c.Request.RequestURI, v) {
				return
			}
		}

		// 状态码
		code := c.Writer.Status()

		// 请求IP
		ip := utils.ClientIP(c)
		refer := c.Request.Header.Get("referer")

		logger.WithFields(logrus.Fields{
			"status": code,
			"method": c.Request.Method,
			"url":    c.Request.URL.String(),
			"ip":     ip,
			"ua":     c.Request.UserAgent(),
			"d":      time.Since(start),
			"refer":  refer,
		}).Info()
	}
}
