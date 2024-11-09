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
		AllowedOrigins string
		ReceiptScannerURL string
	}
)

func New() (*Container, error) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	http := &HTTP{
		Port: os.Getenv("HTTP_PORT"),
		AllowedOrigins: os.Getenv("HTTP_ALLOWED_ORIGINS"),
		ReceiptScannerURL: os.Getenv("RECEIPT_SCANNER_URL"),
	}

	return &Container{
		http,
	}, nil
}