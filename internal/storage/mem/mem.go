package mem

import (
	"github.com/IKostarev/yandex-go-dev/internal/utils"
)

type Mem struct {
	memory map[string]string
}

func (m *Mem) Save(long string) (string, error) {
	short := utils.RandomString()

	m.memory[short] = long

	return short, nil
}

func (m *Mem) Get(short string) string {
	mini := m.memory[short]
	if mini == "" {
		return ""
	}

	return mini
}

func (m *Mem) Close() error {
	return m.Close()
}
