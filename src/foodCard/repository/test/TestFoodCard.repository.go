package testFoodCardRepository

import (
	"errors"
	foodCardDomain "github.com/shinYeongHyeon/eveningIsFood-api/src/foodCard/domain"
	"log"
)

type testDB struct{}

// TestFoodCardRepository : [test] FoodCardRepository
type TestFoodCardRepository struct {
	Repository testDB
}

var testFoodCardRepository *TestFoodCardRepository
var wantCreateFail bool

func init() {
	log.Println("TestFoodCardRepository init")
	testFoodCardRepository = &TestFoodCardRepository{}
}

// GetRepository : [test] Get foodCardRepository
func GetRepository(isCreateFail bool) *TestFoodCardRepository {
	wantCreateFail = isCreateFail
	return testFoodCardRepository
}

// Create : [test] create FoodCard Row
func (foodCardRepository *TestFoodCardRepository) Create(_ foodCardDomain.FoodCard) error {
	if wantCreateFail {
		return errors.New("error occur")
	}

	return nil
}
