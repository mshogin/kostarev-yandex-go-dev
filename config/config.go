package config

import (
	"flag"
	"strings"
)

type Config struct {
	Host         string
	Port         string
	BaseShortURL string
}

func LoadConfig() Config {
	var httpAddr string
	var baseShortURL string

	flag.StringVar(&httpAddr, "a", "http://localhost:8080", "HTTP server address")
	flag.StringVar(&baseShortURL, "b", "http://localhost:8080", "Base shortened URL")
	flag.Parse()

	host, port := splitHostURL(httpAddr)

	return Config{
		Host:         host,
		Port:         port,
		BaseShortURL: baseShortURL,
	}
}

func splitHostURL(httpAddr string) (string, string) {
	value := strings.Split(httpAddr, "://")
	url := strings.Split(value[1], ":")

	return url[0], ":" + url[1]
}
