package foodShopCardModule

import (
	"github.com/gofiber/fiber/v2"
	foodShopCardCommandController "github.com/shinYeongHyeon/eveningIsFood-api/src/foodShopCard/controller/command"
	foodShopCardEntity "github.com/shinYeongHyeon/eveningIsFood-api/src/foodShopCard/entities"
	corePostgres "github.com/shinYeongHyeon/eveningIsFood-api/src/shared/core/postgres"
	"log"
)

// CreateModule : returned fiber.App for foodShopCardModule mounting
func CreateModule() *fiber.App {
	foodShopCardModule := fiber.New()

	migrate()

	foodShopCardModule.Post("/", foodShopCardCommandController.Create)

	return foodShopCardModule
}

func migrate() {
	manager := corePostgres.GetManager()

	err := manager.Db.Table("food_shop_cards").AutoMigrate(&foodShopCardEntity.FoodShopCardEntity{})

	if err != nil {
		log.Fatal("Failed for foodShopCard migrate")
	}

	log.Println("SUCCESS for foodShopCard migrate")
}
