package config

import (
	"flag"
	"os"
	"strings"
)

var (
	HTTPAddr     *string
	BaseShortURL *string
)

func init() {
	if envRunAddr := os.Getenv("SERVER_ADDRESS"); envRunAddr != "" {
		HTTPAddr = &envRunAddr
	}

	if envBaseAddr := os.Getenv("BASE_URL"); envBaseAddr != "" {
		BaseShortURL = &envBaseAddr
	}

	HTTPAddr = flag.String("a", "localhost:8080", "HTTP server address")
	BaseShortURL = flag.String("b", "http://localhost", "Base shortened URL")	
}

func LoadConfig() (string, string) {
	host, port := splitHostURL(*HTTPAddr)

	return host, port
}

func splitHostURL(httpAddr string) (string, string) {
	url := strings.Split(httpAddr, ":")

	return url[0], ":" + url[1]
}
