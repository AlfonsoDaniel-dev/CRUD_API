package handlers

import (
	"CRUD_Task_API/models"
	"github.com/labstack/echo/v4"
	"net/http"
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
