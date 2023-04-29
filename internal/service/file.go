package service

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

type FileStorage struct {
	WriteJSON map[string]string
	file      *os.File
	writer    *bufio.Writer
}

func NewFileStorage(file *os.File) (*FileStorage, error) {
	store := &FileStorage{
		WriteJSON: make(map[string]string),
		file:      file,
		writer:    bufio.NewWriter(file),
	}

	reader := bufio.NewReader(file)

	for {
		id, err := reader.ReadBytes(',')
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		url, err := reader.ReadBytes('\n')
		if err != nil {
			return nil, err
		}

		sID := strings.Trim(string(id), ",")
		sURL := strings.Trim(string(url), "\n")

		store.WriteJSON[sID] = sURL
	}

	return store, nil
}

func (store *FileStorage) Add(url string) (id string, err error) {
	for ok := true; ok; _, ok = store.WriteJSON[id] {
		id = randomString()
	}

	store.WriteJSON[id] = url

	data := []byte(id + "," + url + "\n")
	if _, err := store.writer.Write(data); err != nil {
		return "", err
	}

	if err = store.writer.Flush(); err != nil {
		return "", err
	}

	return id, nil
}

func (store *FileStorage) Get(id string) (string, error) {
	url, ok := store.WriteJSON[id]
	if ok {
		return url, nil
	}
	return "", errors.New("URL not found")
}

func (store *FileStorage) Close() error {
	return store.file.Close()
}
