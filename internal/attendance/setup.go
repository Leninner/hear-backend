package attendance

import (
	"github.com/gofiber/fiber/v2"
	attendanceApp "github.com/leninner/hear-backend/internal/attendance/application"
	attendanceRepo "github.com/leninner/hear-backend/internal/attendance/infrastructure"
	"github.com/leninner/hear-backend/internal/auth/application"
	"github.com/leninner/hear-backend/internal/infrastructure/db"
	sharedMiddleware "github.com/leninner/hear-backend/internal/shared/middleware"
)

func Setup(api fiber.Router, db *db.Queries, authHandler *application.Handler) *attendanceApp.Handler {
	authMiddleware := sharedMiddleware.NewAuthMiddleware(authHandler.GetUseCase().ValidateAccessToken)
	
	// Teacher and student routes
	attendance := api.Group("/attendance", authMiddleware.Authenticate(), authMiddleware.RequireTeacherOrStudent())

	attendanceRepository := attendanceRepo.NewPostgresRepository(db)
	attendanceUseCase := attendanceApp.NewUseCase(attendanceRepository)
	attendanceHandler := attendanceApp.NewHandler(attendanceUseCase)
	attendanceApp.SetupRoutes(attendance, attendanceHandler)

	return attendanceHandler
} 