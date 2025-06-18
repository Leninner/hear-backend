package attendance

import (
	"github.com/gofiber/fiber/v2"
	attendanceApp "github.com/leninner/hear-backend/internal/attendance/application"
	attendanceRepo "github.com/leninner/hear-backend/internal/attendance/infrastructure"
	"github.com/leninner/hear-backend/internal/infrastructure/db"
)

func Setup(api fiber.Router, db *db.Queries) *attendanceApp.Handler {
	attendance := api.Group("/attendance")

	attendanceRepository := attendanceRepo.NewPostgresRepository(db)
	attendanceUseCase := attendanceApp.NewUseCase(attendanceRepository)
	attendanceHandler := attendanceApp.NewHandler(attendanceUseCase)
	attendanceApp.SetupRoutes(attendance, attendanceHandler)

	return attendanceHandler
} 