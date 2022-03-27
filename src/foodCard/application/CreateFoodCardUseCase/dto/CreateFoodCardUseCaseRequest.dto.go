package createFoodCardUseCaseDto

import foodCardDomain "github.com/shinYeongHyeon/eveningIsFood-api/src/foodCard/domain"

// CreateFoodCardUseCaseRequest : Request object for CreateFoodCardUseCase
type CreateFoodCardUseCaseRequest struct {
	Address foodCardDomain.Address
	Name    foodCardDomain.Name
}
