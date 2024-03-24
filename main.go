package main

import (
	"CRUD_Task_API/DB"

	_ "github.com/lib/pq"
)

func main() {
	DB.NewPostgresDB()
}
