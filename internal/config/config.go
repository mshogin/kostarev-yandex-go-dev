package config

import (
	"flag"
	"os"
)

type Config struct {
	ServerAddr   *string
	BaseShortURL *string
}

func LoadConfig() Config {
	serverAddrDefault := ":8080"
	baseShortURL := "http://localhost:8080"

	cfg := Config{
		ServerAddr:   &serverAddrDefault,
		BaseShortURL: &baseShortURL,
	}

	cfg.loadEnv()
	cfg.loadFlags()

	return cfg
}

func (cfg *Config) loadEnv() {
	if envRunAddr := os.Getenv("SERVER_ADDRESS"); envRunAddr != "" {
		cfg.ServerAddr = &envRunAddr
	}

	if envBaseAddr := os.Getenv("BASE_URL"); envBaseAddr != "" {
		cfg.BaseShortURL = &envBaseAddr
	}
}

func (cfg *Config) loadFlags() {
	flag.StringVar(cfg.ServerAddr, "a", *cfg.ServerAddr, "HTTP server address")
	flag.StringVar(cfg.BaseShortURL, "b", *cfg.BaseShortURL, "Base shortened URL")

	flag.Parse()
}
