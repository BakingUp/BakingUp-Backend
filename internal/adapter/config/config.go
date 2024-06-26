package config

import (
	"os"
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
	http := &HTTP{
		Port: os.Getenv("HTTP_PORT"),
	}

	return &Container{
		http,
	}, nil
}