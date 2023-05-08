package fs

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/IKostarev/yandex-go-dev/internal/logger"
	"github.com/IKostarev/yandex-go-dev/internal/utils"
	"io"
	"os"
	"strings"
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

func NewFs(file *os.File) (*Fs, error) {
	fs := &Fs{
		fh:    file,
		cache: make(map[string]string),
		count: 0,
	}

	reader := bufio.NewReader(file)

	for {
		bytes, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, logger.Errorf("reader is error: %w", err)
		}

		line := strings.Trim(string(bytes), "\n")
		spl := strings.Split(line, ",")

		id := spl[0]
		url := spl[1]

		fs.cache[id] = url
	}

	return fs, nil
}

func (m *Fs) Save(long string) (string, error) {
	urlData := &URLData{
		UUID:        fmt.Sprintf("%d", m.count),
		ShortURL:    utils.RandomString(),
		OriginalURL: long,
	}

	jsonData, err := json.Marshal(urlData)
	if err != nil {
		return "", logger.Errorf("cannot marshal json: %w", err)
	}

	_, err = m.fh.Write(jsonData)
	if err != nil {
		return "", logger.Errorf("cannot write to file: %w", err)
	}

	m.count++
	m.cache[urlData.ShortURL] = urlData.OriginalURL

	return urlData.ShortURL, nil
}

func (m *Fs) Get(short string) string {
	return m.cache[short]
}

func (m *Fs) Close() error {
	return m.fh.Close()
}
