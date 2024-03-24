package handlers

import "CRUD_Task_API/models"

type Service interface {
	Create(task *models.Task) error
	GetAll() (models.Tasks, error)
	GetById(id int) (models.Task, error)
	//Update(id int, m *models.Task) error
	//Delete(id int) error
}
