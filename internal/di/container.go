package di

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	attendance "github.com/leninner/hear-backend/internal/attendance"
	attendanceApp "github.com/leninner/hear-backend/internal/attendance/application"
	classroom "github.com/leninner/hear-backend/internal/classroom"
	classroomApp "github.com/leninner/hear-backend/internal/classroom/application"
	courses "github.com/leninner/hear-backend/internal/courses"
	coursesApp "github.com/leninner/hear-backend/internal/courses/application"
	qrcode "github.com/leninner/hear-backend/internal/qrcode"
	qrcodeApp "github.com/leninner/hear-backend/internal/qrcode/application"
	docsApp "github.com/leninner/hear-backend/internal/shared/handler"
	users "github.com/leninner/hear-backend/internal/users"
	userApp "github.com/leninner/hear-backend/internal/users/application"
)

type Container struct {
	App             *fiber.App
	UserHandler     *userApp.Handler
	DocsHandler     *docsApp.DocsHandler
	AttendanceHandler *attendanceApp.Handler
	ClassroomHandler *classroomApp.Handler
	QRCodeHandler   *qrcodeApp.Handler
	CoursesHandler  *coursesApp.Handler
}

func NewContainer(db *sql.DB) *Container {
	app := fiber.New(fiber.Config{
		AppName: "Hear Backend",
	})
	
	// SHARED MODULES
	docsHandler := docsApp.NewDocsHandler()
	docsHandler.Register(app)

	api := app.Group("/api")
	
	// APP MODULES
	userHandler := users.Setup(api, db)
	attendanceHandler := attendance.Setup(api, db)
	classroomHandler := classroom.Setup(api, db)
	qrcodeHandler := qrcode.Setup(api, db)
	coursesHandler := courses.Setup(api, db)

	return &Container{
		App:             app,
		UserHandler:     userHandler,
		DocsHandler:     docsHandler,
		AttendanceHandler: attendanceHandler,
		ClassroomHandler: classroomHandler,
		QRCodeHandler:   qrcodeHandler,
		CoursesHandler:  coursesHandler,
	}
} 