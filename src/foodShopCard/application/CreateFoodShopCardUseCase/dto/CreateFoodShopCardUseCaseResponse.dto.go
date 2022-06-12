package createFoodShopCardUseCaseDto

import foodShopCardDomain "github.com/shinYeongHyeon/eveningIsFood-api/src/foodShopCard/domain"

// CodeFailedSaveFoodShopCardDomain : Failed save foodShopCard
var CodeFailedSaveFoodShopCardDomain = "FAILED_SAVE_FOOD_SHOP_CARD_DOMAIN"

// CreateFoodShopCardUseCaseResponse : Response object for CreateFoodShopCardUseCase
type CreateFoodShopCardUseCaseResponse struct {
	Code         string
	FoodShopCard foodShopCardDomain.FoodShopCard
}
