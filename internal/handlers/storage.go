package handlers

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
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

func (a *App) FileStorage(shortURL, originURL string) error {
	path := a.FilePath()
	if path == "" {
		return errors.New("error path file")
	}

	urlData := &URLData{
		UUID:        fmt.Sprintf("%d", countLines(path)+1),
		ShortURL:    shortURL,
		OriginalURL: originURL,
	}

	jsonData, err := json.Marshal(urlData)
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}

func countLines(filePath string) int {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return 0
	}
	defer file.Close()

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
