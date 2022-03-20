package foodCardRepository

import (
	foodCardDomain "github.com/shinYeongHyeon/eveningIsFoodApi/src/foodCard/domain"
	foodCardEntity "github.com/shinYeongHyeon/eveningIsFoodApi/src/foodCard/entities"
	corePostgres "github.com/shinYeongHyeon/eveningIsFoodApi/src/shared/core/postgres"
	"gorm.io/gorm"
	"log"
)

// FoodCardRepository : struct for foodCardRepository
// TODO: generic 만들고, 모킹 파일 만들어서 테스트시에 그것 활용해보는 것 리서치/Create 호출자도 수정해야 함
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
