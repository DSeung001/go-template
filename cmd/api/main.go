package main

import (
	"context"
	"errors"
	"go-template/internal/handler"
	"go-template/internal/middleware"
	"go-template/pkg/db"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Infra
	mysqlConn, err := db.NewMySQL("root:root@tcp(localhost:3306)/template_db?parseTime=true")
	if err != nil {
		log.Fatalf("failed to connect to mysql: %v", err)
	}
	defer func() {
		if err := mysqlConn.Close(); err != nil {
			log.Printf("failed to close mysql: %v", err)
		}
	}()
	redisClient := db.NewRedis("localhost:6379")
	defer func() {
		if err := redisClient.Close(); err != nil {
			log.Printf("failed to close redis: %v", err)
		}
	}()

	// Engine
	r := gin.Default()

	// MiddleWare
	r.Use(gin.Recovery())
	r.Use(middleware.RequestIDMiddleware())
	r.Use(middleware.RequestLogger())

	// Route
	{
		health := r.Group("/health")
		health.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"message": "ok"}) })
		health.GET("/ready", handler.Health(mysqlConn, redisClient))
	}

	// Server
	srv := &http.Server{
		Addr:    "localhost:8080",
		Handler: r,
	}

	if err := runServer(srv, 5*time.Second); err != nil {
		log.Printf("server exited with error: %v", err)
		os.Exit(1)
	}
	log.Println("server exited")
	os.Exit(0)
}

func runServer(srv *http.Server, shutdownTimeout time.Duration) error {
	errCh := make(chan error, 1)
	go func() {
		errCh <- srv.ListenAndServe()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errCh:
		return listenResult(err)
	case <-quit:
	}

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		return err
	}
	return listenResult(<-errCh)
}

func listenResult(err error) error {
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}
