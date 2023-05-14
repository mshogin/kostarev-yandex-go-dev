package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"

	"github.com/IKostarev/yandex-go-dev/internal/utils"
)

type DB struct {
	db    *pgxpool.Pool
	cache map[string]string
	count int64
}

func NewDB(addrConn string) (*DB, error) {
	addr, err := pgxpool.ParseConfig(addrConn)
	if err != nil {
		return nil, fmt.Errorf("error parse config: %w", err)
	}

	conn, err := pgxpool.NewWithConfig(context.Background(), addr)
	if err != nil {
		return nil, fmt.Errorf("error create NewWithConfig: %w", err)
	}

	db := &DB{
		db:    conn,
		cache: make(map[string]string),
		count: 1,
	}

	exists, err := db.checkIsTablesExists()
	if err != nil {
		return nil, fmt.Errorf("error check is table exists: %w", err)
	}

	if !exists {
		if err = db.createTable(); err != nil {
			return nil, fmt.Errorf("error create tables: %w", err)
		}
	}

	return db, nil
}

func (db *DB) Save(longURL string) (string, error) {
	shortURL := utils.RandomString()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	_, err := db.db.Exec(ctx, "INSERT INTO yandex (uuid, longurl, shorturl) VALUES ($1, $2, $3);", db.count, longURL, shortURL)
	if err != nil {
		return "", fmt.Errorf("error is INSERT data in database: %w", err)
	}

	db.cache[shortURL] = longURL
	db.count++

	return shortURL, nil
}

func (db *DB) Get(shortURL string) string {
	//if longURL, ok := db.cache[shortURL]; ok {
	//	return longURL
	//}
	//
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	//defer cancel()
	//
	//row := db.db.QueryRow(ctx, "SELECT longurl FROM yandex WHERE shorturl = $1", shortURL)
	//
	//var longURL string
	//
	//err := row.Scan(&longURL)
	//if err != nil {
	//	logger.Errorf("error is Scan data in SELECT Query: %s", err)
	//	return ""
	//}
	//
	//db.cache[shortURL] = longURL
	//
	return shortURL // longURL
}

func (db *DB) Close() error {
	return nil //TODO заглушка на будущее, кажется что в бд этот метод вообще не нужен
}

func (db *DB) createTable() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	_, err := db.db.Exec(ctx, "CREATE TABLE yandex (uuid UUID NOT NULL UNIQUE, longurl VARCHAR(2048) NOT NULL, shorturl VARCHAR(64) NOT NULL)")
	if err != nil {
		return fmt.Errorf("error create table in create table: %w", err)
	}

	return nil
}

func (db *DB) checkIsTablesExists() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	row := db.db.QueryRow(ctx, "SELECT EXISTS (SELECT FROM yandex)")

	var res bool

	err := row.Scan(&res)
	if err != nil {
		return false, err
	}

	return res, nil
}
