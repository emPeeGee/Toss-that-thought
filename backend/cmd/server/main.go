package main

import (
	"context"
	"github.com/emPeeee/ttt/internal/config"
	"github.com/emPeeee/ttt/internal/connection"
	"github.com/emPeeee/ttt/internal/cors"
	"github.com/emPeeee/ttt/internal/entity"
	"github.com/emPeeee/ttt/internal/flaw"
	"github.com/emPeeee/ttt/internal/thought"
	"github.com/emPeeee/ttt/pkg/accesslog"
	"github.com/emPeeee/ttt/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const Version = "1.0.0"

// RUN: Before autoMigrate -> CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
func main() {
	logger := log.New().With(nil, "version", Version)

	if err := os.Setenv("TZ", "Universal"); err != nil {
		logger.Fatalf("Error setting environment variable")
	}

	cfg, err := config.Get(logger)
	if err != nil {
		logger.Fatalf("failed to initialize config: %s", err.Error())
	}

	db, err := connection.NewPostgresDB(cfg.DB)
	if err != nil {
		logger.Fatalf("failed to initialize db: %s", err.Error())
	}

	err = db.AutoMigrate(&entity.Thought{})
	if err != nil {
		logger.Fatalf("failed to auto migrate gorm", err.Error())
	}

	server := new(connection.Server)
	valid := validator.New()

	go func() {
		if err := server.Run(cfg.Server, buildHandler(db, valid, logger)); err != nil {
			logger.Fatalf("Error occurred while running http server: %s", err.Error())
		}
	}()

	logger.Info("TTt Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logger.Info("TTt Shutting Down")

	if err := server.Shutdown(context.Background()); err != nil {
		logger.Fatalf("error occurred on server shutting down: %s", err.Error())
	}
}

// buildHandler sets up the HTTP routing and builds an HTTP handler.
func buildHandler(db *gorm.DB, valid *validator.Validate, logger log.Logger) http.Handler {
	router := gin.New()
	router.Use(accesslog.Handler(logger), flaw.Handler(logger), cors.Handler())

	thought.RegisterHandlers(
		router.Group("/api"),
		thought.NewThoughtService(thought.NewRepository(db, logger), logger),
		valid,
		logger,
	)

	return router
}
