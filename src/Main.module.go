package mainModule

import (
	"github.com/gofiber/fiber/v2"
	foodCardModule "github.com/shinYeongHyeon/eveningIsFoodApi/src/foodCard"
	healthModule "github.com/shinYeongHyeon/eveningIsFoodApi/src/health"
	userModule "github.com/shinYeongHyeon/eveningIsFoodApi/src/user"
)

// CreateModule is returned fiber.App for mounting
func CreateModule() *fiber.App {
	mainModule := fiber.New()

	mainModule.Mount("/health", healthModule.CreateModule())
	mainModule.Mount("/user", userModule.CreateModule())
	mainModule.Mount("/foodCard", foodCardModule.CreateModule())

	return mainModule
}
