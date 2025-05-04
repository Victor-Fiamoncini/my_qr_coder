package main

import (
	"log"
	"os"

	"github.com/Victor-Fiamoncini/my_qr_coder/internal/app/service"
	"github.com/Victor-Fiamoncini/my_qr_coder/internal/infra"
	"github.com/Victor-Fiamoncini/my_qr_coder/internal/web"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	skip2QrCodeGenerator := infra.NewSkip2QrCodeGenerator()
	s3FileStorage, err := infra.NewS3FileStorage(os.Getenv("AWS_BUCKET_NAME"), os.Getenv("AWS_REGION"))

	if err != nil {
		panic(err)
	}

	generateQrCodeService := service.NewGenerateQrCodeService(skip2QrCodeGenerator, s3FileStorage)

	httpServer := web.NewHttpServer("8080", generateQrCodeService)

	httpServer.RegisterRoutes()

	if err := httpServer.Start(); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
