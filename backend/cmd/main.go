package main

import (
	"context"
	"github.com/emPeeee/ttt/internal/thought"
	"github.com/emPeeee/ttt/pkg/log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/emPeeee/ttt"
	"github.com/emPeeee/ttt/pkg/repository"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var Version = "1.0.0"

// Here is what from max and that friend
// Max, graceful shutdown, configs,
// Friend, handler, service, repo, logger

// To continue implementing new arhitecture,
// DB init to be removed from repository
// Entity as well ???
// configure logger as in example
// ...
func main() {
	logger := log.New().With(nil, "version", Version)

	if err := initializeConfig(); err != nil {
		logger.Fatalf("Error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logger.Fatalf("Error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
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

	server := new(ttt.Server)

	go func() {
		if err := server.Run(viper.GetString("port"), buildHandler(db, logger)); err != nil {
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
func buildHandler(db *sqlx.DB, logger log.Logger) http.Handler {
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.Use(gin.Logger())

	valid := validator.New()
	thought.RegisterHandlers(router.Group("/api"), thought.NewThoughtService(thought.NewRepository(db, logger), logger), valid, logger)

	return router

}
