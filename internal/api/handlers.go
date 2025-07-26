package api

import "github.com/gofiber/fiber/v2"

func uploadHandler(c *fiber.Ctx) error {
	// parse file → enqueue job → return JSON{job_id}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"job_id": "TODO"})
}

func statusHandler(c *fiber.Ctx) error {
	// fetch job status from DB
	return c.JSON(fiber.Map{"status": "TODO"})
}

func downloadHandler(c *fiber.Ctx) error {
	// stream processed image if done
	return c.SendStatus(fiber.StatusNotImplemented)
}
