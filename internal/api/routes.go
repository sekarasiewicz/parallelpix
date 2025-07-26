package api

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {
	v1 := app.Group("/v1")
	v1.Post("/upload", uploadHandler)
	v1.Get("/status/:id", statusHandler)
	v1.Get("/download/:id", downloadHandler)
}
