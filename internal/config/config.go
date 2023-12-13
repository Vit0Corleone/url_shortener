package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Env string
	HTTPServer
}

type HTTPServer struct {
	Address     string
	Timeout     time.Duration
	IdleTimeout time.Duration
	User        string
	Password    string
}

// panic if error
func MustLoad() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(".env not loaded")
	}

	var cfg = Config{
		os.Getenv("CONFIG_ENV"),
		HTTPServer{
			Address:     os.Getenv("CONFIG_ADDRESS"),
			Timeout:     time.Second * 4,
			IdleTimeout: time.Second * 60,
			User:        os.Getenv("HTTPSERVER_USER"),
			Password:    os.Getenv("HTTPSERVER_PASSWORD"),
		},
	}

	return &cfg
}
