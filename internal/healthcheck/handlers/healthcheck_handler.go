package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Alive(fiberContext *fiber.Ctx) error {
	jsonResponse := map[string]bool{
		"alive": true,
	}
	return fiberContext.Status(200).JSON(jsonResponse)
}
