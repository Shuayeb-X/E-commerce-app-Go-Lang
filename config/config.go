package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int64
}

var configurations Config

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load .env file", err)
		os.Exit(1)
	}

	servicename := os.Getenv("SERVICE_NAME")
	if servicename == "" {
		fmt.Println("Service Name is required")
		os.Exit(1)
	}

	version := os.Getenv("VERSION")

	if version == "" {
		fmt.Println("Version is required")
		os.Exit(1)
	}
	httpPort := os.Getenv("HTTP_PORT")
	port, err := strconv.ParseInt(httpPort, 10, 64)

	if err != nil {
		fmt.Println("Port must be number")
		os.Exit(1)
	}

	configurations = Config{
		Version:     version,
		ServiceName: servicename,
		HttpPort:    port,
	}
}

func GetConfig() Config {
	loadConfig()
	return configurations
}

