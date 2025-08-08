package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DB struct {
	db *sql.DB
}

func New() (*DB, error) {
	db, err := sql.Open("postgres", "postgresql://root:password@localhost:5433/go-chat?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return &DB{
		db: db,
	}, nil
}

func (d *DB) GetDB() *sql.DB {
	return d.db
}

func (d *DB) Close() {
	d.db.Close()
}
