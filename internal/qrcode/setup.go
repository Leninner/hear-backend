package qrcode

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leninner/hear-backend/internal/infrastructure/db"
	qrcodeApp "github.com/leninner/hear-backend/internal/qrcode/application"
	qrcodeRepo "github.com/leninner/hear-backend/internal/qrcode/infrastructure"
)

func Setup(api fiber.Router, db *db.Queries) *qrcodeApp.Handler {
	qrcodes := api.Group("/qrcodes")

	qrcodeRepository := qrcodeRepo.NewPostgresRepository(db)
	qrcodeUseCase := qrcodeApp.NewUseCase(qrcodeRepository)
	qrcodeHandler := qrcodeApp.NewHandler(qrcodeUseCase)
	qrcodeApp.SetupRoutes(qrcodes, qrcodeHandler)

	return qrcodeHandler
} 