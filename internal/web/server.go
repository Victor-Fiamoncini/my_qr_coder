package web

import "github.com/gofiber/fiber/v2"

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
	h.server.Post("/qrcode", func(c *fiber.Ctx) error {
		var body QrCodePostRequestBody

		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON format"})
		}

		if body.Text == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "The 'text' field is required"})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Received text successfully",
			"text":    body.Text,
		})
	})
}

func (h *HttpServer) Start() error {
	h.server.Listen(":" + h.port)

	return nil
}
