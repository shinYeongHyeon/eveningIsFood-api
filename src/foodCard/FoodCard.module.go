package foodCardModule

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	foodCardEntity "github.com/shinYeongHyeon/eveningIsFoodApi/src/foodCard/entities"
	corePostgres "github.com/shinYeongHyeon/eveningIsFoodApi/src/shared/core/postgres"
	"log"
)

// CreateModule : returned fiber.App for foodCardModule mounting
func CreateModule() *fiber.App {
	userModule := fiber.New()

	userModule.Get("/*", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("✋ %s, in FoodCard Module", c.Params("*"))
		return c.SendString(msg)
	})

	migrate()

	return userModule
}

func migrate() {
	manager := corePostgres.GetManager()

	err := manager.Db.Table("food_cards").AutoMigrate(&foodCardEntity.FoodCardEntity{})

	if err != nil {
		log.Fatal("Failed for foodCard migrate")
	}

	log.Println("SUCCESS for foodCard migrate")
}
