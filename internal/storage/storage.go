package storage

type Storage interface {
	Save(string) (string, error)
	Get(string) string
	Close() error
}
