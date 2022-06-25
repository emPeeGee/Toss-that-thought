package main

import (
	"context"
	"github.com/emPeeee/ttt/internal/config"
	"github.com/emPeeee/ttt/internal/connection"
	"github.com/emPeeee/ttt/internal/cors"
	"github.com/emPeeee/ttt/internal/flaw"
	"github.com/emPeeee/ttt/internal/thought"
	"github.com/emPeeee/ttt/pkg/accesslog"
	"github.com/emPeeee/ttt/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const Version = "1.0.0"

// Here is what from max and that friend
// Max, graceful shutdown, configs,
// Friend, handler, service, repo, logger
// My, connection, crypt

func main() {
	logger := log.New().With(nil, "version", Version)
	config, err := config.Get(logger)
	if err != nil {
		logger.Fatalf("failed to initialize config: %s", err.Error())
	}

	db, err := connection.NewPostgresDB(config.DB)

	if err != nil {
		logger.Fatalf("failed to initialize db: %s", err.Error())
	}

	server := new(connection.Server)
	valid := validator.New()

	go func() {
		if err := server.Run(config.Server, buildHandler(db, valid, logger)); err != nil {
			logger.Fatalf("Error occurred while running http server: %s", err.Error())
		}
	}()

	logger.Info("TTt Started")
	port, err := net.LookupPort("tcp", "http")
	logger.Info(port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logger.Info("TTt Shutting Down")

	if err := server.Shutdown(context.Background()); err != nil {
		logger.Fatalf("error occurred on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logger.Fatalf("error occurred on db connection close: %s", err.Error())
	}
}

// buildHandler sets up the HTTP routing and builds an HTTP handler.
func buildHandler(db *sqlx.DB, valid *validator.Validate, logger log.Logger) http.Handler {
	router := gin.New()
	router.Use(accesslog.Handler(logger), flaw.Handler(logger), cors.Handler())

	thought.RegisterHandlers(router.Group("/api"), thought.NewThoughtService(thought.NewRepository(db, logger), logger), valid, logger)

	return router
}
