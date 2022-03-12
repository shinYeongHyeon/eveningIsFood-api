package foodCardModule

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

// CreateModule : returned fiber.App for foodCardModule mounting
func CreateModule() *fiber.App {
	userModule := fiber.New()

	userModule.Get("/*", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("✋ %s, in FoodCard Module", c.Params("*"))
		return c.SendString(msg)
	})

	return userModule
}
