package main

import (
	migrations "CRUD_Task_API/Migrations"
	"CRUD_Task_API/db"
	"CRUD_Task_API/routes"
	"fmt"
	"github.com/labstack/echo/v4"

	_ "github.com/lib/pq"
)

func main() {
	DB.NewPostgresDB()

	if err := migrations.NewPsqlProductMigration(DB.Pool()); err != nil {
		fmt.Printf("Error al correr la migracion. Err:%v", err)
	}

	e := echo.New()
	db := DB.Pool()

	TaskStorage := DB.PsqlTaskConstructor(db)

	routes.TasksRoutes(e, TaskStorage)

	if err := e.Start("192.168.5.207:8000"); err != nil {
		fmt.Printf("Error al iniciar el servidor: %v", err)
	}

}
