package postgres

import (
	"context"
	"fmt"
	"github.com/IKostarev/yandex-go-dev/internal/logger"
	"github.com/IKostarev/yandex-go-dev/internal/utils"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type DB struct {
	db    *pgxpool.Pool
	cache map[string]string
	count int64
}

func NewDB(addrConn string) (*DB, error) {
	addr, err := pgxpool.ParseConfig(addrConn)
	if err != nil {
		logger.Errorf("error parse config: %s", err)
	}

	conn, err := pgxpool.NewWithConfig(context.Background(), addr)
	if err != nil {
		logger.Errorf("error create NewWithConfig: %s", err)
	}

	db := &DB{
		db:    conn,
		cache: make(map[string]string),
		count: 1,
	}

	if err = db.createTable(); err != nil {
		logger.Errorf("error create tables: %s", err)
	}

	return db, nil
}

func (db *DB) Save(longURL string) (string, error) {
	shortURL := utils.RandomString()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	_, err := db.db.Exec(ctx, "INSERT INTO yandex_url (id, long_url, long_url_id) VALUES ($1, $2, $3);", db.count, longURL, shortURL)
	if err != nil {
		return "", fmt.Errorf("error is INSERT data in database: %w", err)
	}

	db.cache[shortURL] = longURL
	db.count++

	return shortURL, nil
}

func (db *DB) Get(shortURL string) string {
	var longURL string

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	row, err := db.db.Query(ctx, "SELECT long_url FROM yandex_url WHERE long_url_id = $1", shortURL)
	if err != nil {
		logger.Errorf("error is SELECT data in database: %s", err)
		return ""
	}

	err = row.Scan(&longURL)
	if err != nil {
		logger.Errorf("error is Scan data in SELECT Query: %s", err)
		return ""
	}

	return longURL
}

func (db *DB) Close() error {
	return nil //TODO заглушка на будущее, кажется что в бд этот метод вообще не нужен
}

func (db *DB) createTable() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	_, err := db.db.Exec(ctx, "CREATE TABLE yandex_url (id VARCHAR(255) NOT NULL UNIQUE, long_url VARCHAR(255) NOT NULL, long_url_id uuid NOT NULL )")
	if err != nil {
		return err
	}

	return nil
}
