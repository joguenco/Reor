package handlers

import "github.com/gofiber/fiber/v2"

type Ping struct {
	Message string `json:"message"`
}

func GetPing(c *fiber.Ctx) error {
	return c.JSON(Ping{Message: "pong"})
}
