package config

import (
	"flag"
	"log"
	"net/url"
	"os"
	"strings"
)

var (
	HTTPAddr     *string
	BaseShortURL *string
)

type Config struct {
	Host    string
	Port    string
	BaseURL string
}

func LoadConfig() Config {
	if envRunAddr := os.Getenv("SERVER_ADDRESS"); envRunAddr != "" {
		HTTPAddr = &envRunAddr
	} else {
		HTTPAddr = flag.String("a", "localhost:8080", "HTTP server address")
	}

	if envBaseAddr := os.Getenv("BASE_URL"); envBaseAddr != "" {
		BaseShortURL = &envBaseAddr
	} else {
		BaseShortURL = flag.String("b", "http://localhost", "Base shortened URL")
	}

	if _, err := url.ParseRequestURI(*BaseShortURL); err != nil {
		log.Fatal("you didn't enter a url")
	}

	host, port := splitHostURL(*HTTPAddr)

	return Config{
		Host:    host,
		Port:    port,
		BaseURL: *BaseShortURL,
	}
}

func splitHostURL(httpAddr string) (string, string) {
	url := strings.Split(httpAddr, ":")

	return url[0], ":" + url[1]
}
