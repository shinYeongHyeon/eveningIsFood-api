package mainModule

import (
	"github.com/gofiber/fiber/v2"
	foodCardModule "github.com/shinYeongHyeon/eveningIsFoodApi/src/foodCard"
	healthModule "github.com/shinYeongHyeon/eveningIsFoodApi/src/health"
	corePostgres "github.com/shinYeongHyeon/eveningIsFoodApi/src/shared/core/postgres"
	userModule "github.com/shinYeongHyeon/eveningIsFoodApi/src/user"
)

// CreateModule is returned fiber.App for mounting
func CreateModule() *fiber.App {
	mainModule := fiber.New()

	mainModule.Mount("/health", healthModule.CreateModule())
	mainModule.Mount("/user", userModule.CreateModule())
	mainModule.Mount("/foodCard", foodCardModule.CreateModule())

	// NOTE: 연결 테스트
	corePostgres.GetManager()

	return mainModule
}
