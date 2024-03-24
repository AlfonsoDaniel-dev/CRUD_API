package DB

import (
	"CRUD_Task_API/models"
	"database/sql"
	"fmt"
)

const (
	SqlCreateTask  = `INSERT INTO tasks(id, title, content, created_at) VALUES ($1, $2, $3, $4)`
	SqlGetAllTasks = `SELECT *  FROM tasks`
	SqlGetTaskById = `SELECT (id, title, content, created_at) FROM tasks WHERE id = $1`
)

type PsqlTask struct {
	database *sql.DB
}

func PsqlTaskConstructor(db *sql.DB) *PsqlTask {
	return &PsqlTask{
		database: db,
	}
}

func (t *PsqlTask) Create(m *models.Task) error {
	stmt, err := t.database.Prepare(SqlCreateTask)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		&m.ID,
		&m.Title,
		&m.Content,
		&m.Created_at,
	)

	if err != nil {
		return err
	}

	fmt.Println("Task Creada correctamente")

	return nil
}

func (t *PsqlTask) GetAll() (models.Tasks, error) {
	stmt, err := t.database.Prepare(SqlGetAllTasks)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()

	tasks := make(models.Tasks, 0)
	for rows.Next() {
		T := &models.Task{}
		err = rows.Scan(&T.ID, &T.Title, &T.Content, &T.Created_at)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		tasks = append(tasks, T)
	}

	return tasks, nil
}

func (T *PsqlTask) GetById(id int) (models.Task, error) {
	stmt, err := T.database.Prepare(SqlGetTaskById)
	if err != nil {
		return models.Task{}, err
	}
	defer stmt.Close()

	var Task models.Task
	row := stmt.QueryRow(id)
	err = row.Scan(&Task.ID, &Task.Title, &Task.Content, &Task.Created_at)
	if err != nil {
		return models.Task{}, err
	}

	return Task, nil
}
