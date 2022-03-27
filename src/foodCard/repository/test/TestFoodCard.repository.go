package testFoodCardRepository

import (
	foodCardDomain "github.com/shinYeongHyeon/eveningIsFood-api/src/foodCard/domain"
	"log"
)

type testDB struct{}

// TestFoodCardRepository : [test] FoodCardRepository
type TestFoodCardRepository struct {
	Repository testDB
}

var testFoodCardRepository *TestFoodCardRepository

func init() {
	log.Println("TestFoodCardRepository init")
	testFoodCardRepository = &TestFoodCardRepository{}
}

// GetRepository : [test] Get foodCardRepository
func GetRepository() *TestFoodCardRepository {
	return testFoodCardRepository
}

// Create : [test] create FoodCard Row
func (foodCardRepository *TestFoodCardRepository) Create(_ foodCardDomain.FoodCard) error {
	return nil
}
