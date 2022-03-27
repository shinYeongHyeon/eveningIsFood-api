package foodCardRepository

import (
	foodCardDomain "github.com/shinYeongHyeon/eveningIsFood-api/src/foodCard/domain"
	postgresFoodCardRepository "github.com/shinYeongHyeon/eveningIsFood-api/src/foodCard/repository/postgres"
	testFoodCardRepository "github.com/shinYeongHyeon/eveningIsFood-api/src/foodCard/repository/test"
)

// IFoodCardRepository : Generics for repositories
type IFoodCardRepository interface {
	*postgresFoodCardRepository.FoodCardRepository | *testFoodCardRepository.TestFoodCardRepository
	Create(f foodCardDomain.FoodCard) error
}
