package config

import (
	"flag"
	"strings"
)

var (
	HttpAddr     *string
	BaseShortURL *string
)

func init() {
	HttpAddr = flag.String("a", "localhost:8080", "HTTP server address")
	BaseShortURL = flag.String("b", "http://localhost", "Base shortened URL")
}

func LoadConfig() (string, string) {
	host, port := splitHostURL(*HttpAddr)

	return host, port
}

func splitHostURL(httpAddr string) (string, string) {
	url := strings.Split(httpAddr, ":")

	return url[0], ":" + url[1]
}
