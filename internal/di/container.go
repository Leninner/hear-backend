package di

import (
	"github.com/gofiber/fiber/v2"
	attendance "github.com/leninner/hear-backend/internal/attendance"
	attendanceApp "github.com/leninner/hear-backend/internal/attendance/application"
	auth "github.com/leninner/hear-backend/internal/auth"
	authApp "github.com/leninner/hear-backend/internal/auth/application"
	classroom "github.com/leninner/hear-backend/internal/classroom"
	classroomApp "github.com/leninner/hear-backend/internal/classroom/application"
	courses "github.com/leninner/hear-backend/internal/courses"
	coursesApp "github.com/leninner/hear-backend/internal/courses/application"
	"github.com/leninner/hear-backend/internal/infrastructure/db"
	qrcode "github.com/leninner/hear-backend/internal/qrcode"
	qrcodeApp "github.com/leninner/hear-backend/internal/qrcode/application"
	docsApp "github.com/leninner/hear-backend/internal/shared/handler"
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
}

func NewContainer(db *db.Queries) *Container {
	app := fiber.New(fiber.Config{
		AppName: "Hear Backend",
	})
	
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

	return &Container{
		App:             app,
		AuthHandler:     authHandler,
		UserHandler:     userHandler,
		DocsHandler:     docsHandler,
		AttendanceHandler: attendanceHandler,
		ClassroomHandler: classroomHandler,
		QRCodeHandler:   qrcodeHandler,
		CoursesHandler:  coursesHandler,
	}
} 