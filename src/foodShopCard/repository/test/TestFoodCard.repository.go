package testFoodShopCardRepository

import (
	"errors"
	foodShopCardDomain "github.com/shinYeongHyeon/eveningIsFood-api/src/foodShopCard/domain"
	"log"
)

type testDB struct{}

// TestFoodShopCardRepository : [test] FoodShopCardRepository
type TestFoodShopCardRepository struct {
	Repository testDB
}

var testFoodShopCardRepository *TestFoodShopCardRepository
var wantCreateFail bool

func init() {
	log.Println("TestFoodShopCardRepository init")
	testFoodShopCardRepository = &TestFoodShopCardRepository{}
}

// GetRepository : [test] Get foodShopCardRepository
func GetRepository(isCreateFail bool) *TestFoodShopCardRepository {
	wantCreateFail = isCreateFail
	return testFoodShopCardRepository
}

// Create : [test] create FoodShopCard Row
func (foodShopCardRepository *TestFoodShopCardRepository) Create(_ foodShopCardDomain.FoodShopCard) error {
	if wantCreateFail {
		return errors.New("error occur")
	}

	return nil
}
