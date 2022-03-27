package foodCardRepository

import (
	postgresFoodCardRepository "github.com/shinYeongHyeon/eveningIsFood-api/src/foodCard/repository/postgres"
	testFoodCardRepository "github.com/shinYeongHyeon/eveningIsFood-api/src/foodCard/repository/test"
)

// IFoodCardRepository : Generics for repositories
type IFoodCardRepository interface {
	postgresFoodCardRepository.FoodCardRepository | testFoodCardRepository.TestFoodCardRepository
}
