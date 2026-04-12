package middleware

import (
	"go-template/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)

		logger.Log.WithFields(logrus.Fields{
			"status":     c.Writer.Status(),
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"ip":         c.ClientIP(),
			"latency":    latency.Milliseconds(),
			"user_agent": c.Request.UserAgent(),
			"request-id": c.GetString(RequestIDHeader),
		}).Info("request_completed")
	}
}
