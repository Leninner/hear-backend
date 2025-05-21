package main

import (
	"database/sql"
	"log"

	"github.com/leninner/hear-backend/internal/di"
	server "github.com/leninner/hear-backend/internal/shared/server"
	_ "github.com/lib/pq"
)

func main() {
	sqlDB, err := sql.Open("postgres", "postgres://hear:hear@localhost:5432/hear?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	container := di.NewContainer(sqlDB)
	server.SetupServer(container)

	log.Fatal(container.App.Listen(":8080"))
}
