package foodCardModule

import (
	"github.com/gofiber/fiber/v2"
	foodCardCommandController "github.com/shinYeongHyeon/eveningIsFood-api/src/foodCard/controller/command"
	foodCardEntity "github.com/shinYeongHyeon/eveningIsFood-api/src/foodCard/entities"
	corePostgres "github.com/shinYeongHyeon/eveningIsFood-api/src/shared/core/postgres"
	"log"
)

// CreateModule : returned fiber.App for foodCardModule mounting
func CreateModule() *fiber.App {
	foodCardModule := fiber.New()

	migrate()

	foodCardModule.Post("/", foodCardCommandController.Create)

	return foodCardModule
}

func migrate() {
	manager := corePostgres.GetManager()

	err := manager.Db.Table("food_cards").AutoMigrate(&foodCardEntity.FoodCardEntity{})

	if err != nil {
		log.Fatal("Failed for foodCard migrate")
	}

	log.Println("SUCCESS for foodCard migrate")
}
