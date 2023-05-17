package mem

import (
	"github.com/IKostarev/yandex-go-dev/internal/utils"
)

type Mem struct {
	cacheMemory      map[string]string
	cacheCorrelation map[string]string
}

func NewMem() (*Mem, error) {
	m := &Mem{
		cacheMemory:      make(map[string]string),
		cacheCorrelation: make(map[string]string),
	}

	return m, nil
}

func (m *Mem) Save(long, corrId string) (string, error) {
	short := utils.RandomString()

	if long != "" && corrId == "" {
		m.cacheMemory[short] = long
		return short, nil
	}

	m.cacheCorrelation[corrId] = long
	return corrId, nil
}

func (m *Mem) Get(short, corrId string) (string, string) {
	if short != "" && corrId == "" {
		return m.cacheMemory[short], corrId
	}

	return m.cacheCorrelation[corrId], corrId
}

func (m *Mem) Close() error {
	return nil
}
