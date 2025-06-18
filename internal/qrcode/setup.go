package qrcode

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	qrcodeApp "github.com/leninner/hear-backend/internal/qrcode/application"
	qrcodeRepo "github.com/leninner/hear-backend/internal/qrcode/infrastructure"
)

func Setup(api fiber.Router, db *sql.DB) *qrcodeApp.Handler {
	qrcodes := api.Group("/qrcodes")

	qrcodeRepository := qrcodeRepo.NewPostgresRepository(db)
	qrcodeUseCase := qrcodeApp.NewUseCase(qrcodeRepository)
	qrcodeHandler := qrcodeApp.NewHandler(qrcodeUseCase)
	qrcodeApp.SetupRoutes(qrcodes, qrcodeHandler)

	return qrcodeHandler
} 