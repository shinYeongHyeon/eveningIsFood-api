package postgresFoodCardRepository

import (
	foodCardDomain "github.com/shinYeongHyeon/eveningIsFood-api/src/foodCard/domain"
	foodCardEntity "github.com/shinYeongHyeon/eveningIsFood-api/src/foodCard/entities"
	corePostgres "github.com/shinYeongHyeon/eveningIsFood-api/src/shared/core/postgres"
	"gorm.io/gorm"
	"log"
)

// FoodCardRepository : struct for foodCardRepository
type FoodCardRepository struct {
	Repository *gorm.DB
}

var foodCardRepository *FoodCardRepository

func init() {
	log.Println("FoodCardRepository init")
	foodCardRepository = &FoodCardRepository{corePostgres.GetManager().Db.Table("food_cards")}
}

// GetRepository : Get FoodCardRepository
func GetRepository() *FoodCardRepository {
	return foodCardRepository
}

// Create : create FoodCard Row
func (foodCardRepository *FoodCardRepository) Create(foodCard foodCardDomain.FoodCard) error {
	foodCardRepository.Repository.Create(&foodCardEntity.FoodCardEntity{
		UUID:    foodCard.GetUuid(),
		Name:    foodCard.GetName().Value(),
		Address: foodCard.GetAddress().Value(),
	})

	return nil
}
