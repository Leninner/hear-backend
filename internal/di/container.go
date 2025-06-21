package di

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	attendance "github.com/leninner/hear-backend/internal/attendance"
	attendanceApp "github.com/leninner/hear-backend/internal/attendance/application"
	auth "github.com/leninner/hear-backend/internal/auth"
	authApp "github.com/leninner/hear-backend/internal/auth/application"
	classroom "github.com/leninner/hear-backend/internal/classroom"
	classroomApp "github.com/leninner/hear-backend/internal/classroom/application"
	courses "github.com/leninner/hear-backend/internal/courses"
	coursesApp "github.com/leninner/hear-backend/internal/courses/application"
	faculties "github.com/leninner/hear-backend/internal/faculties"
	facultiesApp "github.com/leninner/hear-backend/internal/faculties/application"
	"github.com/leninner/hear-backend/internal/infrastructure/db"
	qrcode "github.com/leninner/hear-backend/internal/qrcode"
	qrcodeApp "github.com/leninner/hear-backend/internal/qrcode/application"
	docsApp "github.com/leninner/hear-backend/internal/shared/handler"
	universities "github.com/leninner/hear-backend/internal/universities"
	universitiesApp "github.com/leninner/hear-backend/internal/universities/application"
	users "github.com/leninner/hear-backend/internal/users"
	userApp "github.com/leninner/hear-backend/internal/users/application"
)

type Container struct {
	App             *fiber.App
	AuthHandler     *authApp.Handler
	UserHandler     *userApp.Handler
	DocsHandler     *docsApp.DocsHandler
	AttendanceHandler *attendanceApp.Handler
	ClassroomHandler *classroomApp.Handler
	QRCodeHandler   *qrcodeApp.Handler
	CoursesHandler  *coursesApp.Handler
	UniversitiesHandler *universitiesApp.Handler
	FacultiesHandler *facultiesApp.Handler
}

func NewContainer(db *db.Queries) *Container {
	app := fiber.New(fiber.Config{
		AppName: "Hear Backend",
	})
	
	// Configure CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "*",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization,X-Requested-With",
		ExposeHeaders:    "*",
	}))
	
	// SHARED MODULES
	docsHandler := docsApp.NewDocsHandler()
	docsHandler.Register(app)

	api := app.Group("/api")
	
	// AUTH MODULE (no auth required)
	authHandler := auth.Setup(api, db)
	
	// APP MODULES (each handles its own middleware)
	userHandler := users.Setup(api, db, authHandler)
	attendanceHandler := attendance.Setup(api, db, authHandler)
	classroomHandler := classroom.Setup(api, db, authHandler)
	qrcodeHandler := qrcode.Setup(api, db, authHandler)
	coursesHandler := courses.Setup(api, db, authHandler)
	universitiesHandler := universities.Setup(api, db, authHandler)
	facultiesHandler := faculties.Setup(api, db, authHandler)

	return &Container{
		App:             app,
		AuthHandler:     authHandler,
		UserHandler:     userHandler,
		DocsHandler:     docsHandler,
		AttendanceHandler: attendanceHandler,
		ClassroomHandler: classroomHandler,
		QRCodeHandler:   qrcodeHandler,
		CoursesHandler:  coursesHandler,
		UniversitiesHandler: universitiesHandler,
		FacultiesHandler: facultiesHandler,
	}
} 