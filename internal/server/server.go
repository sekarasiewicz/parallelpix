package server

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sekarasiewicz/parallelpix/internal/api"
)

func Start() {
	app := fiber.New()

	// mount all API routes
	api.RegisterRoutes(app)

	data, _ := json.MarshalIndent(app.Stack(), "", "  ")
	fmt.Print(string(data))

	// health check, metrics, pprof etc. could go here too

	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
