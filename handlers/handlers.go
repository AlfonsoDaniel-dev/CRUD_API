package handlers

import (
	"CRUD_Task_API/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type HandlerService struct {
	service Service
}

func NewHandlerService(s Service) *HandlerService {
	return &HandlerService{
		service: s,
	}
}

func (Hs *HandlerService) Create(c echo.Context) error {
	data := models.Task{}
	err := c.Bind(&data)
	if err != nil {
		response := NewResponse(Error, "Peticion mal estructurada", err)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = Hs.service.Create(&data)
	if err != nil {
		response := NewResponse(Error, "Error al Crear La Tarea", err)
		return c.JSON(http.StatusBadRequest, response)
	}

	response := NewResponse(MessageOK, "tasks Creada correctamente", data)
	return c.JSON(http.StatusCreated, response)
}

func (Hs *HandlerService) GetAll(c echo.Context) error {
	tasks, err := Hs.service.GetAll()
	if err != nil {
		response := NewResponse(Error, "No se pudo obtener las tareas", err)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := NewResponse(MessageOK, "Tasks obtenidas correctamente", tasks)
	return c.JSON(http.StatusOK, response)
}

func (Hs *HandlerService) GetById(c echo.Context) error {
	IDstr := c.Param("id")

	ID, err := strconv.Atoi(IDstr)
	if err != nil {
		response := NewResponse(Error, "Id invalido", err)
		return c.JSON(http.StatusBadRequest, response)
	}

	Task, err := Hs.service.GetById(ID)
	if err != nil {
		response := NewResponse(Error, "No se pudo obtener la tares", err)
		return c.JSON(http.StatusNotFound, response)
	}

	response := NewResponse(MessageOK, "Tarea obtenida correctamente", Task)
	return c.JSON(http.StatusOK, response)
}

func (Hs *HandlerService) Update(c echo.Context) error {
	IdStr := c.Param("id")

	ID, err := strconv.Atoi(IdStr)
	if err != nil {
		response := NewResponse(Error, "Id no valido", err)
		return c.JSON(http.StatusBadRequest, response)
	}

	task := models.Task{}

	err = c.Bind(&task)
	if err != nil {
		response := NewResponse(Error, "Peticion mal estrucutrada", err)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = Hs.service.Update(ID, &task)
	if err != nil {
		response := NewResponse(Error, "No se pudo actualizar la tarea", err)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := NewResponse(MessageOK, "Tarea actualizada correctamente", task)
	return c.JSON(http.StatusOK, response)

}

func (Hs *HandlerService) Delete(c echo.Context) error {
	IdStr := c.Param("id")

	ID, err := strconv.Atoi(IdStr)
	if err != nil {
		response := NewResponse(Error, "id no valido", err)
		return c.JSON(http.StatusInternalServerError, response)
	}

	err = Hs.service.Delete(ID)
	if err != nil {
		response := NewResponse(Error, "No se pudo eliminar la tarea", err)
		return c.JSON(http.StatusBadRequest, response)
	}

	response := NewResponse(MessageOK, "Se Elimino la tarea correctamente", nil)
	return c.JSON(http.StatusOK, response)
}
