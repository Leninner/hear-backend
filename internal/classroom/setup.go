package classroom

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	classroomApp "github.com/leninner/hear-backend/internal/classroom/application"
	classroomRepo "github.com/leninner/hear-backend/internal/classroom/infrastructure"
)

func Setup(api fiber.Router, db *sql.DB) *classroomApp.Handler {
	classrooms := api.Group("/classrooms")

	classroomRepository := classroomRepo.NewPostgresRepository(db)
	classroomUseCase := classroomApp.NewUseCase(classroomRepository)
	classroomHandler := classroomApp.NewHandler(classroomUseCase)
	classroomApp.SetupRoutes(classrooms, classroomHandler)

	return classroomHandler
} 