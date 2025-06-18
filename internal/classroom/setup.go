package classroom

import (
	"github.com/gofiber/fiber/v2"
	classroomApp "github.com/leninner/hear-backend/internal/classroom/application"
	classroomRepo "github.com/leninner/hear-backend/internal/classroom/infrastructure"
	"github.com/leninner/hear-backend/internal/infrastructure/db"
)

func Setup(api fiber.Router, db *db.Queries) *classroomApp.Handler {
	classrooms := api.Group("/classrooms")

	classroomRepository := classroomRepo.NewPostgresRepository(db)
	classroomUseCase := classroomApp.NewUseCase(classroomRepository)
	classroomHandler := classroomApp.NewHandler(classroomUseCase)
	classroomApp.SetupRoutes(classrooms, classroomHandler)

	return classroomHandler
} 