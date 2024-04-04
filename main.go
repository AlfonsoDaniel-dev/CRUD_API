package main

import (
	"CRUD_Task_API/DB"
	migrations "CRUD_Task_API/Migrations"
	"CRUD_Task_API/routes"
	"fmt"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	DB.NewPostgresDB()

	migrator := migrations.NewMigration(DB.Pool())
	if err := migrator.MigrateUser(); err != nil {
		log.Fatal(err)
	}

	if err := migrator.MigrateTask(); err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	db := DB.Pool()

	TaskStorage := DB.PsqlTaskConstructor(db)

	routes.TasksRoutes(e, TaskStorage)

	if err := e.Start(":8000"); err != nil {
		fmt.Printf("Error al iniciar el servidor: %v", err)
	}

}
