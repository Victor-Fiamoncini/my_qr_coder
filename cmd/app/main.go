package main

import "github.com/Victor-Fiamoncini/my_qr_coder/internal/web"

func main() {
	httpServer := web.NewHttpServer("8080")

	httpServer.RegisterRoutes()
	httpServer.Start()
}
