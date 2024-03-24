package db

import (
	"database/sql"
	"sync"
)

var (
	db   *sql.DB
	once sync.Once
)

func NewPostgresDB() {
	once.Do(func() {

	})
}
