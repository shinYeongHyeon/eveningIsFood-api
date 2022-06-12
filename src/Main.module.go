package mainModule

import (
	"github.com/gofiber/fiber/v2"
	foodShopCardModule "github.com/shinYeongHyeon/eveningIsFood-api/src/foodShopCard"
	healthModule "github.com/shinYeongHyeon/eveningIsFood-api/src/health"
	userModule "github.com/shinYeongHyeon/eveningIsFood-api/src/user"
)

// CreateModule is returned fiber.App for mounting
func CreateModule() *fiber.App {
	mainModule := fiber.New()

	mainModule.Mount("/health", healthModule.CreateModule())
	mainModule.Mount("/user", userModule.CreateModule())
	mainModule.Mount("/foodShopCard", foodShopCardModule.CreateModule())

	return mainModule
}
