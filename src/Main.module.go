package mainModule

import (
	"github.com/gofiber/fiber/v2"
	foodCardModule "github.com/shinYeongHyeon/eveningIsFood-api/src/foodCard"
	healthModule "github.com/shinYeongHyeon/eveningIsFood-api/src/health"
	oauthModule "github.com/shinYeongHyeon/eveningIsFood-api/src/oauth"
	userModule "github.com/shinYeongHyeon/eveningIsFood-api/src/user"
)

// CreateModule is returned fiber.App for mounting
func CreateModule() *fiber.App {
	mainModule := fiber.New()

	mainModule.Mount("/health", healthModule.CreateModule())
	mainModule.Mount("/user", userModule.CreateModule())
	mainModule.Mount("/foodCard", foodCardModule.CreateModule())
	mainModule.Mount("/oauth", oauthModule.CreateModule())

	return mainModule
}
