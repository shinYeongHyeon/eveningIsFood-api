package userModule

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	corePostgres "github.com/shinYeongHyeon/eveningIsFoodApi/src/shared/core/postgres"
	userEntity "github.com/shinYeongHyeon/eveningIsFoodApi/src/user/entities"
)

// CreateModule : returned fiber.App for UserModule mounting
func CreateModule() *fiber.App {
	userModule := fiber.New()

	userModule.Get("/*", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("✋ %s, in User Module", c.Params("*"))
		return c.SendString(msg)
	})

	migrate()

	return userModule
}

func migrate() {
	manager := corePostgres.GetManager()

	err := manager.Db.Table("users").AutoMigrate(&userEntity.UserEntity{})

	if err != nil {
		log.Fatal("Failed for user migrate")
	}

	log.Println("SUCCESS for user migrate")
}
