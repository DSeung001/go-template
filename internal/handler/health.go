package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

func Health(mysqlConn *sqlx.DB, redisClient *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
		defer cancel()

		mysqlPingErr := mysqlConn.PingContext(ctx)
		redisPingErr := redisClient.Ping(ctx).Err()

		if mysqlPingErr != nil || redisPingErr != nil {
			response := map[string]any{}

			if mysqlPingErr != nil {
				response["mysql"] = mysqlPingErr.Error()
			}

			if redisPingErr != nil {
				response["redis"] = redisPingErr.Error()
			}

			response["message"] = "unhealthy"

			c.JSON(http.StatusServiceUnavailable, response)
			return
		}

		c.JSON(200, gin.H{"message": "ok"})
	}
}
