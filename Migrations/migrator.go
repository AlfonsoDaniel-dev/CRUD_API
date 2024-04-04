package migrations

import "database/sql"

type Migrator struct {
	dataBase *sql.DB
}

func NewMigration(db *sql.DB) *Migrator {
	return &Migrator{
		dataBase: db,
	}
}
