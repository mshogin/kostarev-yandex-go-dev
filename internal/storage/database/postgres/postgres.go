package postgres

import (
	"context"
	"github.com/IKostarev/yandex-go-dev/internal/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	db *pgxpool.Pool
}

func NewDB(addrConn string) (*DB, error) {
	addr, err := pgxpool.ParseConfig(addrConn)
	if err != nil {
		logger.Errorf("error parse config: %s", err)
	}

	conn, err := pgxpool.NewWithConfig(context.Background(), addr)
	if err != nil {
		logger.Errorf("error create NewWithConfig : %s", err)
	}

	db := &DB{
		db: conn,
	}

	return db, nil
}

func (db *DB) Save(saveStr string) (string, error) {
	return saveStr, nil //TODO заглушка на будущее
}

func (db *DB) Get(getStr string) string {
	return getStr //TODO заглушка на будущее
}

func (db *DB) Close() error {
	return nil //TODO заглушка на будущее
}
