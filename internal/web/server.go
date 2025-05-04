package web

import (
	"os"

	"github.com/Victor-Fiamoncini/my_qr_coder/internal/app/service"
	"github.com/Victor-Fiamoncini/my_qr_coder/internal/infra"
	"github.com/gofiber/fiber/v2"
)

type QrCodePostRequestBody struct {
	Text string `json:"text"`
}

type HttpServer struct {
	server *fiber.App
	port   string
}

func NewHttpServer(port string) *HttpServer {
	return &HttpServer{
		server: fiber.New(),
		port:   port,
	}
}

func (h *HttpServer) RegisterRoutes() {
	skip2QrCodeGenerator := infra.NewSkip2QrCodeGenerator()

	region := os.Getenv("AWS_REGION")
	bucketName := os.Getenv("AWS_BUCKET_NAME")

	s3FileStorage, err := infra.NewS3FileStorage(region, bucketName)

	if err != nil {
		panic(err)
	}

	generateQrCodeService := service.NewGenerateQrCodeService(skip2QrCodeGenerator, s3FileStorage)

	h.server.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("Server is running")
	})

	h.server.Post("/qrcode", func(c *fiber.Ctx) error {
		var body QrCodePostRequestBody

		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON format"})
		}

		if body.Text == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "The 'text' field is required"})
		}

		qrCodeUrl, err := generateQrCodeService.GenerateQrCode(body.Text)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate QR code"})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"url": qrCodeUrl})
	})
}

func (h *HttpServer) Start() error {
	h.server.Listen(":" + h.port)

	return nil
}
