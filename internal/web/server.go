package web

import (
	"fmt"

	"github.com/Victor-Fiamoncini/my_qr_coder/internal/app/service"
	"github.com/gofiber/fiber/v2"
)

type QrCodePostRequestBody struct {
	Text string `json:"text"`
}

type HttpServer struct {
	server                *fiber.App
	port                  string
	generateQrCodeService *service.GenerateQrCodeService
}

func NewHttpServer(port string, generateQrCodeService *service.GenerateQrCodeService) *HttpServer {
	return &HttpServer{
		server:                fiber.New(),
		port:                  port,
		generateQrCodeService: generateQrCodeService,
	}
}

func (h *HttpServer) RegisterRoutes() {
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

		qrCodeUrl, err := h.generateQrCodeService.GenerateQrCode(body.Text)

		fmt.Println("err", err)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate QR code"})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"url": qrCodeUrl})
	})
}

func (h *HttpServer) Start() error {
	err := h.server.Listen(":" + h.port)

	if err != nil {
		return err
	}

	return nil
}
