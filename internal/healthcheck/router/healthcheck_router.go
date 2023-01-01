package router

import (
	"buyer/internal/healthcheck/handlers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(router fiber.Router, db *gorm.DB) {
	router.Get("/alive", func(c *fiber.Ctx) error {
		return handlers.Alive(c)
	})
}
