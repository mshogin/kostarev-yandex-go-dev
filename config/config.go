package config

import (
	"flag"
)

// Config структура конфигурации.
type Config struct {
	HTTPAddr     string // Адрес HTTP-сервера.
	BaseShortURL string // Базовый адрес результирующего сокращённого URL.
}

// LoadConfig инициализирует и возвращает структуру конфигурации.
func LoadConfig() Config {
	var httpAddr string
	var baseShortURL string

	// Обработка флагов командной строки.
	flag.StringVar(&httpAddr, "a", "http://localhost:8080", "HTTP server address")
	flag.StringVar(&baseShortURL, "b", "http://localhost:8080", "Base shortened URL")
	flag.Parse()

	// Проверка наличия корректного протокола в базовом URL.
	//if baseShortURL[:4] != "http" {
	//	baseShortURL = fmt.Sprintf("http://%s", baseShortURL)
	//}

	return Config{
		HTTPAddr:     httpAddr,
		BaseShortURL: baseShortURL,
	}
}
