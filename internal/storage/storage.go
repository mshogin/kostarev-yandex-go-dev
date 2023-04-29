package storage

import (
	"github.com/IKostarev/yandex-go-dev/internal/config"
	"github.com/IKostarev/yandex-go-dev/internal/service"
	"log"
	"os"
)

type Storage interface {
	Add(url string) (id string, err error)
	Get(id string) (string, error)
}

func NewStorage() (Storage, error) {
	var s Storage

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error load config in storage: ", err)
	}

	if path := cfg.FileStoragePath; path != "" {
		var file *os.File

		file, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0777)
		if err != nil {
			log.Fatal("Error open file is: ", err)
		}

		s, err = service.NewFileStorage(file)
	}

	if err != nil {
		return nil, err
	}

	return s, nil
}
