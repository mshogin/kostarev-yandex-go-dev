package mem

import (
	"fmt"
	"github.com/IKostarev/yandex-go-dev/internal/utils"
)

type Mem struct {
	memory map[string]string
}

func NewMem() (*Mem, error) {
	m := &Mem{
		memory: make(map[string]string),
	}

	return m, nil
}

func (m *Mem) Save(long string) (string, error) {
	short := utils.RandomString()

	fmt.Println("save long = ", long)
	fmt.Println("save short = ", short)

	m.memory[short] = long

	fmt.Println("save m.memory[short] = ", m.memory[short])

	return short, nil
}

func (m *Mem) Get(short string) string {
	mini := m.memory[short]
	if mini == "" {
		return ""
	}

	fmt.Println("get mem short = ", short)
	fmt.Println("get mem mini = ", mini)

	return mini
}

func (m *Mem) Close() error {
	return nil
}
