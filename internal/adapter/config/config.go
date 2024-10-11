package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type (
	Container struct {
		HTTP *HTTP
		DB   *DB
	}

	DB struct {
		Connection string
		Host       string
		Port       string
		User       string
		Password   string
		Name       string
	}

	HTTP struct {
		URL  string
		Port string
	}
)

func New() (*Container, error) {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error while reading env file")
		return nil, err
	}

	HTTP := &HTTP{
		URL:  os.Getenv("HTTP_URL"),
		Port: os.Getenv("HTTP_PORT"),
	}

	DB := &DB{
		Connection: os.Getenv("DB_CONNECTION"),
		Host:       os.Getenv("DB_HOST"),
		Port:       os.Getenv("DB_PORT"),
		User:       os.Getenv("DB_USER"),
		Password:   os.Getenv("DB_PASSWORD"),
		Name:       os.Getenv("DB_NAME"),
	}

	return &Container{
		HTTP: HTTP,
		DB:   DB,
	}, nil
}
