package migrations

import (
	"database/sql"
	"fmt"
)

const SqlMigrateTask = `CREATE TABLE IF NOT EXISTS tasks(
    id SERIAL NOT NULL,
    title varchar(75) NOT NULL,
    content varchar(240) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    CONSTRAINT tasks_id_pk PRIMARY KEY (id)
)`

type Migrator struct {
	dataBase *sql.DB
}

func NewPsqlProductMigration(db *sql.DB) *Migrator {
	return &Migrator{
		dataBase: db,
	}
}

func (m *Migrator) Migrate() error {
	stmt, err := m.dataBase.Prepare(SqlMigrateTask)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("Migracion de task realizada")

	return nil

}
