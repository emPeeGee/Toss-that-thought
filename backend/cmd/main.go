package main

import (
	"context"
	"github.com/emPeeee/ttt/internal/connection"
	"github.com/emPeeee/ttt/internal/cors"
	"github.com/emPeeee/ttt/internal/flaw"
	"github.com/emPeeee/ttt/internal/thought"
	"github.com/emPeeee/ttt/pkg/accesslog"
	"github.com/emPeeee/ttt/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var Version = "1.0.0"

// Here is what from max and that friend
// Max, graceful shutdown, configs,
// Friend, handler, service, repo, logger

func main() {
	logger := log.New().With(nil, "version", Version)

	if err := initializeConfig(); err != nil {
		logger.Fatalf("Error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logger.Fatalf("Error loading env variables: %s", err.Error())
	}

	db, err := connection.NewPostgresDB(connection.DBConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logger.Fatalf("failed to initialize db: %s", err.Error())
	}

	server := new(connection.Server)
	valid := validator.New()

	go func() {
		if err := server.Run(viper.GetString("port"), buildHandler(db, valid, logger)); err != nil {
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

	if err := db.Close(); err != nil {
		logger.Fatalf("error occurred on db connection close: %s", err.Error())
	}
}

func initializeConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

// buildHandler sets up the HTTP routing and builds an HTTP handler.
func buildHandler(db *sqlx.DB, valid *validator.Validate, logger log.Logger) http.Handler {
	router := gin.New()
	router.Use(accesslog.Handler(logger), flaw.Handler(logger), cors.Handler())

	thought.RegisterHandlers(router.Group("/api"), thought.NewThoughtService(thought.NewRepository(db, logger), logger), valid, logger)

	return router
}
