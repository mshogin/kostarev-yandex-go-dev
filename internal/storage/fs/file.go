package fs

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/IKostarev/yandex-go-dev/internal/config"
	"github.com/IKostarev/yandex-go-dev/internal/logger"
	"github.com/IKostarev/yandex-go-dev/internal/utils"
	"os"
)

type Fs struct {
	fh    *os.File
	cache map[string]string
	count int64
}

type URLData struct {
	UUID        string `json:"uuid"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

func Load(cfg config.Config) *Fs {
	file, _ := os.OpenFile(cfg.FileStoragePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	//TODO не знаю как обработать здесь ошибку
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			logger.Error("close file is error: ", err)
		}
	}(file)

	reader := bufio.NewReader(file)

	var count int64
	for {
		_, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		count++
	}

	return &Fs{
		fh:    file,
		count: count,
	}
}

func (m *Fs) Save(long string) (string, error) {
	urlData := &URLData{
		UUID:        fmt.Sprintf("%d", m.count),
		ShortURL:    utils.RandomString(),
		OriginalURL: long,
	}

	jsonData, err := json.Marshal(urlData)
	if err != nil {
		return "cannot marshal json", err
	}

	_, err = m.fh.Write(jsonData)
	if err != nil {
		return "cannot write to file", err
	}

	m.count++

	return urlData.ShortURL, nil
}

func (m *Fs) Get(short string) string {
	return m.cache[short]
}

func (m *Fs) Close() error {
	return m.fh.Close()
}
