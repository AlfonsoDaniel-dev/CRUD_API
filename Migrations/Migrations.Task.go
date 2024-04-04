package migrations

import (
	"fmt"
)

const SqlMigrateTask = `CREATE TABLE IF NOT EXISTS tasks(
    id SERIAL NOT NULL,
    user_id SERIAL NOT NULL,
    title VARCHAR(45) NOT NULL,
    content VARCHAR(240) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    CONSTRAINT tasks_id_pk PRIMARY KEY (id),
    CONSTRAINT tasks_user_id_fk FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
)
`

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

	fmt.Println("Task Migration Success")

	return nil

}
