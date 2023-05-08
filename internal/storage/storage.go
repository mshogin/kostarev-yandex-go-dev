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

func NewStorage(cfg config.Config) (s Storage, err error) {
	if path := cfg.FileStoragePath; path != "" {
		file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
		if err != nil {
			return nil, logger.Errorf("cannot open file: %w", err)
		}

		s, err = fs.NewFs(file)
	} else {
		s, err = mem.NewMem()
	}
	if err != nil {
		return nil, logger.Errorf("cannot create storage: %w", err)
	}

	return s, nil
}
