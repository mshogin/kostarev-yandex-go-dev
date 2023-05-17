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

func (m *Mem) Save(long, corrID string) (string, error) {
	short := utils.RandomString()

	if long != "" && corrID == "" {
		m.cacheMemory[short] = long
		return short, nil
	}

	m.cacheCorrelation[corrID] = long
	return corrID, nil
}

func (m *Mem) Get(short, corrID string) (string, string) {
	if short != "" && corrID == "" {
		return m.cacheMemory[short], corrID
	}

	return m.cacheCorrelation[corrID], corrID
}

func (m *Mem) Close() error {
	return nil
}
