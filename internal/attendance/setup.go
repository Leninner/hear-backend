package attendance

import (
	"github.com/gofiber/fiber/v2"
	attendanceApp "github.com/leninner/hear-backend/internal/attendance/application"
	attendanceRepo "github.com/leninner/hear-backend/internal/attendance/infrastructure"
	"github.com/leninner/hear-backend/internal/auth/application"
	classroomRepo "github.com/leninner/hear-backend/internal/classroom/infrastructure"
	"github.com/leninner/hear-backend/internal/infrastructure/db"
	schedulesRepo "github.com/leninner/hear-backend/internal/schedules/infrastructure"
	sharedMiddleware "github.com/leninner/hear-backend/internal/shared/middleware"
)

func Setup(api fiber.Router, db *db.Queries, authHandler *application.Handler) *attendanceApp.Handler {
	authMiddleware := sharedMiddleware.NewAuthMiddleware(authHandler.GetUseCase().ValidateAccessToken)
	
	// Teacher and student routes
	attendance := api.Group("/attendance", authMiddleware.Authenticate(), authMiddleware.RequireAnyRole("admin", "teacher", "student"))

	attendanceRepository := attendanceRepo.NewPostgresRepository(db)
	scheduleRepository := schedulesRepo.NewPostgresRepository(db)
	classroomRepository := classroomRepo.NewPostgresRepository(db)
	
	// Create adapters for the attendance domain
	scheduleAdapter := attendanceRepo.NewScheduleAdapter(scheduleRepository)
	classroomAdapter := attendanceRepo.NewClassroomAdapter(classroomRepository)
	
	attendanceUseCase := attendanceApp.NewUseCase(attendanceRepository, scheduleAdapter, classroomAdapter)
	attendanceHandler := attendanceApp.NewHandler(attendanceUseCase)
	attendanceApp.SetupRoutes(attendance, attendanceHandler)

	return attendanceHandler
} 