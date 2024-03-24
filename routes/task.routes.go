package routes

import (
	"CRUD_Task_API/handlers"
	"CRUD_Task_API/middlewares"
	"github.com/labstack/echo/v4"
)

func TasksRoutes(e *echo.Echo, s handlers.Service) {
	handler := handlers.NewHandlerService(s)
	tasks := e.Group("/API/v1/tasks", middlewares.Log)

	tasks.GET("", handler.GetAll)
	tasks.POST("/", handler.Create)
}
