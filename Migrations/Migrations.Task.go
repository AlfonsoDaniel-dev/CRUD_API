package migrations

import (
	"database/sql"
	"fmt"
)

const SqlMigrateTask = `ALTER TABLE tasks ADD user_id INTEGER FOREIGN KEY (user_id) REFERENCES users(id);
)`

type Migrator struct {
	dataBase *sql.DB
}

func NewPsqlProductMigration(db *sql.DB) *Migrator {
	return &Migrator{
		dataBase: db,
	}
}

func (m *Migrator) MigrateTask() error {
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
