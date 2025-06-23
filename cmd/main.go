package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/leninner/hear-backend/internal/di"
	"github.com/leninner/hear-backend/internal/infrastructure/db"
	_ "github.com/lib/pq"
)

func main() {
	sqlDB := setupDB()
	defer sqlDB.Close()

	container := di.NewContainer(sqlDB)
	container.App.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	log.Fatal(container.App.Listen(":8080"))
}

func setupDB() *db.Queries {
	dbHost := flag.String("db-host", "localhost", "Database host")
	dbPort := flag.String("db-port", "5432", "Database port")
	dbUser := flag.String("db-user", "hear", "Database user")
	dbPassword := flag.String("db-password", "hear", "Database password")
	dbName := flag.String("db-name", "hear", "Database name")
	flag.Parse()

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		*dbUser, *dbPassword, *dbHost, *dbPort, *dbName)

	sqlDB, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = sqlDB.Ping(); err != nil {
		log.Fatal(err)
	}

	return db.New(sqlDB)
}
