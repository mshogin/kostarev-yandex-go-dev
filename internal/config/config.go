package config

import (
	"flag"
	"log"
	"net/url"
	"os"
)

const (
	serverAddrDefault   = ":8080"
	baseShortURLDefault = "http://localhost:8080"
)

type Config struct {
	ServerAddr   string
	BaseShortURL string
}

func LoadConfig() Config {
	cfg := Config{
		ServerAddr:   serverAddrDefault,
		BaseShortURL: baseShortURLDefault,
	}

	cfg.loadEnv()
	cfg.loadFlags()
	cfg.validate()

	return cfg
}

func (cfg *Config) loadEnv() {
	if envRunAddr := os.Getenv("SERVER_ADDRESS"); envRunAddr != "" {
		cfg.ServerAddr = envRunAddr
	}

	if envBaseAddr := os.Getenv("BASE_URL"); envBaseAddr != "" {
		cfg.BaseShortURL = envBaseAddr
	}
}

func (cfg *Config) loadFlags() {
	flag.StringVar(&cfg.ServerAddr, "a", cfg.ServerAddr, "HTTP server address")
	flag.StringVar(&cfg.BaseShortURL, "b", cfg.BaseShortURL, "Base shortened URL")

	flag.Parse()
}

func (cfg *Config) validate() {
	_, err := url.Parse(cfg.BaseShortURL)
	if err != nil {
		log.Fatal("You entered an invalid URL")
	}
}
