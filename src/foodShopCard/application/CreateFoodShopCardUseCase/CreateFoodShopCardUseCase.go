package createFoodShopCardUseCase

import (
	createFoodShopCardUseCaseDto "github.com/shinYeongHyeon/eveningIsFood-api/src/foodShopCard/application/CreateFoodShopCardUseCase/dto"
	foodShopCardDomain "github.com/shinYeongHyeon/eveningIsFood-api/src/foodShopCard/domain"
	foodShopCardRepository "github.com/shinYeongHyeon/eveningIsFood-api/src/foodShopCard/repository"
	coreUseCaseResponseCode "github.com/shinYeongHyeon/eveningIsFood-api/src/shared/core/useCase"
)

// Execute : Create FoodShopCard domain and save it repository
func Execute[T foodShopCardRepository.IFoodShopCardRepository](
	request createFoodShopCardUseCaseDto.CreateFoodShopCardUseCaseRequest,
	repository T,
) createFoodShopCardUseCaseDto.CreateFoodShopCardUseCaseResponse {
	foodShopCard, _ := foodShopCardDomain.CreateNewFoodShopCard(foodShopCardDomain.NewFoodShopCardProps{
		Name:    request.Name,
		Address: request.Address,
	})

	saveError := repository.Create(foodShopCard)

	if saveError != nil {
		return createFoodShopCardUseCaseDto.CreateFoodShopCardUseCaseResponse{
			Code: createFoodShopCardUseCaseDto.CodeFailedSaveFoodShopCardDomain,
		}
	}

	return createFoodShopCardUseCaseDto.CreateFoodShopCardUseCaseResponse{
		Code:         coreUseCaseResponseCode.SUCCESS,
		FoodShopCard: foodShopCard,
	}
}
