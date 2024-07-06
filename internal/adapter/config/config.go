package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type (
	Container struct {
		HTTP *HTTP
	}

	HTTP struct {
		Port string
	}
)

func New() (*Container, error) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	http := &HTTP{
		Port: os.Getenv("HTTP_PORT"),
	}

	return &Container{
		http,
	}, nil
}