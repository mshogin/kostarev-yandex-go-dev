package handlers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/IKostarev/yandex-go-dev/internal/logger"
	"os"
)

type URLData struct {
	UUID        string `json:"uuid"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

func (a *App) FilePath() string {
	return a.Config.FileStoragePath
}

func (a *App) FileStorage(shortURL, originURL string) {
	path := a.FilePath()
	if path == "" {
		return
	}

	urlData := &URLData{
		UUID:        fmt.Sprintf("%d", countLines(path)+1),
		ShortURL:    shortURL,
		OriginalURL: originURL,
	}

	jsonData, err := json.Marshal(urlData)
	if err != nil {
		return
	}

	file, err := os.Create(path)
	if err != nil {
		return
	}
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			logger.Error("close file is error: ", err)
		}
	}(file)

	_, err = file.Write(jsonData)
	if err != nil {
		return
	}
}

func countLines(filePath string) int {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return 0
	}
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			logger.Error("close file is error: ", err)
		}
	}(file)

	reader := bufio.NewReader(file)

	count := 0
	for {
		_, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		count++
	}

	return count
}
