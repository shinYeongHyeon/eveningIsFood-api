package createFoodCardUseCaseDto

import foodCardDomain "github.com/shinYeongHyeon/eveningIsFood-api/src/foodCard/domain"

// CodeFailedSaveFoodCardDomain : Failed save foodCard
var CodeFailedSaveFoodCardDomain = "FAILED_SAVE_FOOD_CARD_DOMAIN"

// CreateFoodCardUseCaseResponse : Response object for CreateFoodCardUseCase
type CreateFoodCardUseCaseResponse struct {
	Code     string
	FoodCard foodCardDomain.FoodCard
}
