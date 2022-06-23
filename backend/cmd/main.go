package main

import (
	"context"
	"github.com/emPeeee/ttt/pkg/log"
	"github.com/go-playground/validator/v10"
	"os"
	"os/signal"
	"syscall"

	"github.com/emPeeee/ttt"
	"github.com/emPeeee/ttt/pkg/handler"
	"github.com/emPeeee/ttt/pkg/repository"
	"github.com/emPeeee/ttt/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var Version = "1.0.0"

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

	v := validator.New()
	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services, v)

	go func() {
		if err := server.Run(viper.GetString("port"), handlers.InitializeRoutes()); err != nil {
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
