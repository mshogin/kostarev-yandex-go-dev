package config

import (
	"flag"
)

type Config struct {
	HTTPAddr     string
	BaseShortURL string
}

func LoadConfig() *Config {
	httpAddr := flag.String("a", "localhost:8888", "HTTP server address")
	baseShortURL := flag.String("b", "", "base shortened URL")

	flag.Parse()

	return &Config{
		HTTPAddr:     *httpAddr,
		BaseShortURL: *baseShortURL,
	}
}
