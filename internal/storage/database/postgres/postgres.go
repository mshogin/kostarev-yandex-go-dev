package postgres

import (
	"context"
	"fmt"
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

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

	exists, err := db.checkIsTablesExists(ctx)
	if err != nil {
		return nil, fmt.Errorf("error check is table exists: %w", err)
	}

	if !exists {
		if err = db.createTable(ctx); err != nil {
			return nil, fmt.Errorf("error create tables: %w", err)
		}
	}

	return db, nil
}

func (db *DB) Save(longURL string) (string, error) {
	shortURL := utils.RandomString()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	_, err := db.db.Exec(ctx, `INSERT INTO yandex (id, longurl, shorturl) VALUES ($1, $2, $3);`, db.count, longURL, shortURL)
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
	//var longURL string
	//
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	//defer cancel()
	//
	//row, err := db.db.Query(ctx, "SELECT longurl FROM yandex WHERE shorturl = $1", shortURL)
	//if err != nil {
	//	logger.Errorf("error is SELECT data in database: %s", err)
	//	return ""
	//}
	//
	//err = row.Scan(&longURL)
	//if err != nil {
	//	logger.Errorf("error is Scan data in SELECT Query: %s", err)
	//	return ""
	//}
	//
	//return longURL
	return shortURL
}

func (db *DB) Close() error {
	return nil //TODO заглушка на будущее, кажется что в бд этот метод вообще не нужен
}

func (db *DB) createTable(ctx context.Context) error {
	_, err := db.db.Exec(ctx, `CREATE TABLE IF NOT EXISTS yandex (id VARCHAR(255) NOT NULL UNIQUE, longurl VARCHAR(255) NOT NULL, shorturl VARCHAR(255) NOT NULL )`)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) checkIsTablesExists(ctx context.Context) (bool, error) {
	row := db.db.QueryRow(ctx, `SELECT COUNT(*) FROM yandex`)

	var res int

	err := row.Scan(&res)
	if err != nil {
		return false, fmt.Errorf("error check is table exists: %w", err)
	}

	if res > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
