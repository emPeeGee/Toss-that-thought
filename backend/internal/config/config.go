package config

import (
	"github.com/emPeeee/ttt/pkg/log"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
	"time"
)

const (
	defaultMaxHeaderBytes = 1 << 20 // 1 MB
	defaultReadTimeout    = 10 * time.Second
	defaultWriteTimeout   = 10 * time.Second
	path                  = "configs"
	fileName              = "config"
	tokenTTL              = time.Hour * 12
)

type Config struct {
	Server
	DB
	Auth
}

type Server struct {
	Addr           string
	MaxHeaderBytes int
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
}

type DB struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type Auth struct {
	SigningKey string
	TokenTTL   time.Duration
}

func Get(logger log.Logger) (*Config, error) {
	if err := initializeConfig(); err != nil {
		logger.Fatalf("Error initializing config: %s", err.Error())
		return nil, err
	}

	if err := godotenv.Load(); err != nil {
		logger.Fatalf("Error loading env variables: %s", err.Error())
		return nil, err
	}

	db := DB{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	server := Server{
		Addr:           ":" + viper.GetString("server.port"),
		MaxHeaderBytes: defaultMaxHeaderBytes,
		ReadTimeout:    defaultReadTimeout,
		WriteTimeout:   defaultWriteTimeout,
	}

	auth := Auth{
		SigningKey: viper.GetString("auth.signingKey"),
		TokenTTL:   tokenTTL,
	}

	return &Config{server, db, auth}, nil

}

func initializeConfig() error {
	viper.AddConfigPath(path)
	viper.SetConfigName(fileName)

	return viper.ReadInConfig()
}
