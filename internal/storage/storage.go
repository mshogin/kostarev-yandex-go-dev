package storage

import (
	"github.com/IKostarev/yandex-go-dev/internal/config"
	"github.com/IKostarev/yandex-go-dev/internal/logger"
	"github.com/IKostarev/yandex-go-dev/internal/storage/fs"
	"github.com/IKostarev/yandex-go-dev/internal/storage/mem"
	"os"
)

type Storage interface {
	Save(string) (string, error)
	Get(string) string
	Close() error
}

func NewStorage(cfg config.Config) (Storage, error) {
	var s Storage
	var err error

	if path := cfg.FileStoragePath; path != "" {
		var file *os.File

		file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
		if err != nil {
			return nil, err
		}
		defer func() {
			if err = file.Close(); err != nil {
				logger.Error("file close is error in storage: ", err)
			}
		}()

		s, err = fs.NewFs(file)
	} else {
		s, err = mem.NewMem()
	}
	if err != nil {
		return nil, err
	}

	return s, nil
}
