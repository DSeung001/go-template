package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go-template/internal/handler"
	"go-template/internal/middleware"
	"go-template/pkg/db"
	"log"
)

func main() {
	// Infra
	mysqlConn, err := db.NewMySQL("root:root@tcp(localhost:3306)/template_db")
	if err != nil {
		log.Fatalf("failed to connect to mysql: %v", err)
	}
	redisClient := db.NewRedis("localhost:6379")

	// Engine
	r := gin.Default()

	// MiddleWare
	r.Use(gin.Recovery())
	r.Use(middleware.RequestLogger())

	// Route
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})

	r.GET("/health/ready", handler.Health(mysqlConn, redisClient))

	r.GET("/debug/panic", func(c *gin.Context) {
		panic("forced panic for recovery testing")
	})

	if err := r.Run("localhost:8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
