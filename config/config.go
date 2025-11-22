package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}
type ServerConfig struct {
	Port             string
	ReadTimeout      time.Duration
	WriteTimeout     time.Duration
	GracefulShutdown time.Duration
}

func NewConfig() (*Config, error) {
	err := godotenv.Load("settings.env")
	if err != nil {
		return nil, err
	}
	dbuser := os.Getenv("POSTGRES_USER")
	dbpassword := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	dbhost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	serverPort := os.Getenv("SERVER_PORT")
	serverRTimeout := os.Getenv("SERVER_READ_TIMEOUT")
	serverWTimeout := os.Getenv("SERVER_WRITE_TIMEOUT")
	serverGSTimeout := os.Getenv("SERVER_GRACEFUL_SHUTDOWN_TIMEOUT")
	serverR, err := time.ParseDuration(serverRTimeout)
	if err != nil {
		return nil, err
	}
	serverW, err := time.ParseDuration(serverWTimeout)
	if err != nil {
		return nil, err
	}
	serverGS, err := time.ParseDuration(serverGSTimeout)
	if err != nil {
		return nil, err
	}
	dbP, err := strconv.Atoi(dbPort)
	if err != nil {
		return nil, err
	}
	return &Config{
		Database: DatabaseConfig{
			Host:     dbhost,
			Name:     dbname,
			Password: dbpassword,
			User:     dbuser,
			Port:     dbP,
		}, Server: ServerConfig{
			Port:             serverPort,
			GracefulShutdown: serverGS,
			ReadTimeout:      serverR,
			WriteTimeout:     serverW,
		},
	}, nil
}
