package migrations

import (
	"database/sql"
	"fmt"
)

const psqlMigrateUser = `CREATE TABLE IF NOT EXISTS users (
    id SERIAL NOT NULL,
    name varchar(50) NOT NULL,
    email varchar(40) NOT NULL,
    password varchar(255) NOT NULL,
    CONSTRAINT users_id_pk PRIMARY KEY (id),
    CONSTRAINT users_email_uq UNIQUE (email)
);`

type psqlMigrateUser struct {
	db *sql.DB
}

func NewUser(DB *sql.DB) psqlMigrateUser {
	return psqlMigrateUser{DB}
}

func (U *psqlUser) Migrate() error {
	stmt, err := U.db.Prepare(psqlMigrateUser)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("migrated user successfully")

	return nil
}
