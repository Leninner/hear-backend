package qrcode

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leninner/hear-backend/internal/auth/application"
	"github.com/leninner/hear-backend/internal/infrastructure/db"
	qrcodeApp "github.com/leninner/hear-backend/internal/qrcode/application"
	qrcodeRepo "github.com/leninner/hear-backend/internal/qrcode/infrastructure"
	sharedMiddleware "github.com/leninner/hear-backend/internal/shared/middleware"
)

func Setup(api fiber.Router, db *db.Queries, authHandler *application.Handler) *qrcodeApp.Handler {
	authMiddleware := sharedMiddleware.NewAuthMiddleware(authHandler.GetUseCase().ValidateAccessToken)
	
	// Teacher routes only
	qrcodes := api.Group("/qrcodes", authMiddleware.Authenticate(), authMiddleware.RequireTeacher())

	qrcodeRepository := qrcodeRepo.NewPostgresRepository(db)
	qrcodeUseCase := qrcodeApp.NewUseCase(qrcodeRepository)
	qrcodeHandler := qrcodeApp.NewHandler(qrcodeUseCase)
	qrcodeApp.SetupRoutes(qrcodes, qrcodeHandler)

	return qrcodeHandler
} 