package foodShopCardRepository

import (
	foodShopCardDomain "github.com/shinYeongHyeon/eveningIsFood-api/src/foodShopCard/domain"
	postgresFoodShopCardRepository "github.com/shinYeongHyeon/eveningIsFood-api/src/foodShopCard/repository/postgres"
	testFoodShopCardRepository "github.com/shinYeongHyeon/eveningIsFood-api/src/foodShopCard/repository/test"
)

// IFoodShopCardRepository : Generics for repositories
type IFoodShopCardRepository interface {
	*postgresFoodShopCardRepository.FoodShopCardRepository | *testFoodShopCardRepository.TestFoodShopCardRepository
	Create(f foodShopCardDomain.FoodShopCard) error
}
