package main

import (
	migrations "CRUD_Task_API/Migrations"
	"CRUD_Task_API/db"
	"CRUD_Task_API/routes"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	DB.NewPostgresDB()

	migrator := migrations.NewPsqlProductMigration(DB.Pool())
	if err := migrator.Migrate(); err != nil {
		log.Fatalf("No se pudo realizar la migracion de producto. Err: %v", err)
	}

	e := echo.New()
	db := DB.Pool()

	TaskStorage := DB.PsqlTaskConstructor(db)

	routes.TasksRoutes(e, TaskStorage)

	if err := e.Start("192.168.1.2:8000"); err != nil {
		fmt.Printf("Error al iniciar el servidor: %v", err)
	}

}
