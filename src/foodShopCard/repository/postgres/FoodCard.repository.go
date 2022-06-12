package postgresFoodShopCardRepository

import (
	foodShopCardDomain "github.com/shinYeongHyeon/eveningIsFood-api/src/foodShopCard/domain"
	foodShopCardEntity "github.com/shinYeongHyeon/eveningIsFood-api/src/foodShopCard/entities"
	corePostgres "github.com/shinYeongHyeon/eveningIsFood-api/src/shared/core/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

// FoodShopCardRepository : struct for foodShopCardRepository
type FoodShopCardRepository struct {
	Repository *gorm.DB
}

var foodShopCardRepository *FoodShopCardRepository

func init() {
	log.Println("FoodShopCardRepository init")

	if os.Getenv("TESTING") == "true" {
		return
	}

	foodShopCardRepository = &FoodShopCardRepository{corePostgres.GetManager().Db.Table("food_cards")}
}

// GetRepository : Get FoodShopCardRepository
func GetRepository() *FoodShopCardRepository {
	return foodShopCardRepository
}

// Create : create FoodShopCard Row
func (foodShopCardRepository *FoodShopCardRepository) Create(foodShopCard foodShopCardDomain.FoodShopCard) error {
	foodShopCardRepository.Repository.Create(&foodShopCardEntity.FoodShopCardEntity{
		UUID:    foodShopCard.GetUuid(),
		Name:    foodShopCard.GetName().Value(),
		Address: foodShopCard.GetAddress().Value(),
	})

	return nil
}
