package createFoodShopCardUseCaseDto

import foodShopCardDomain "github.com/shinYeongHyeon/eveningIsFood-api/src/foodShopCard/domain"

// CreateFoodShopCardUseCaseRequest : Request object for CreateFoodShopCardUseCase
type CreateFoodShopCardUseCaseRequest struct {
	Address foodShopCardDomain.Address
	Name    foodShopCardDomain.Name
}
