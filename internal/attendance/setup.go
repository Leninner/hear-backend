package attendance

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	attendanceApp "github.com/leninner/hear-backend/internal/attendance/application"
	attendanceRepo "github.com/leninner/hear-backend/internal/attendance/infrastructure"
)

func Setup(api fiber.Router, db *sql.DB) *attendanceApp.Handler {
	attendance := api.Group("/attendance")

	attendanceRepository := attendanceRepo.NewPostgresRepository(db)
	attendanceUseCase := attendanceApp.NewUseCase(attendanceRepository)
	attendanceHandler := attendanceApp.NewHandler(attendanceUseCase)
	attendanceApp.SetupRoutes(attendance, attendanceHandler)

	return attendanceHandler
} 