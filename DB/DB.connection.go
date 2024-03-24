package DB

import (
	"database/sql"
	"log"
	"sync"
)

var (
	db   *sql.DB
	once sync.Once
)

func NewPostgresDB() {
	once.Do(func() {
		var err error
		var connStr = "postgres://poncho:SecurePassword@localhost:8080/CRUDTASKS?sslmode=disable"
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Fatalf("can't connect to DATABASE. Err: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Can't ping DataBase. Err: %v", err)
		}
	})
}

func Pool() *sql.DB {
	return db
}
