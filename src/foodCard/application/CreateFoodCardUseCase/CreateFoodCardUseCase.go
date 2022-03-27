package createFoodCardUseCase

import (
	createFoodCardUseCaseDto "github.com/shinYeongHyeon/eveningIsFood-api/src/foodCard/application/CreateFoodCardUseCase/dto"
	foodCardDomain "github.com/shinYeongHyeon/eveningIsFood-api/src/foodCard/domain"
	foodCardRepository "github.com/shinYeongHyeon/eveningIsFood-api/src/foodCard/repository"
	coreUseCaseResponseCode "github.com/shinYeongHyeon/eveningIsFood-api/src/shared/core/useCase"
)

// Execute : Create FoodCard domain and save it repository
func Execute[T foodCardRepository.IFoodCardRepository](
	request createFoodCardUseCaseDto.CreateFoodCardUseCaseRequest,
	repository T,
) createFoodCardUseCaseDto.CreateFoodCardUseCaseResponse {
	foodCard, _ := foodCardDomain.CreateNewFoodCard(foodCardDomain.NewFoodCardProps{
		Name:    request.Name,
		Address: request.Address,
	})

	saveError := repository.Create(foodCard)

	if saveError != nil {
		return createFoodCardUseCaseDto.CreateFoodCardUseCaseResponse{
			Code: createFoodCardUseCaseDto.CodeFailedSaveFoodCardDomain,
		}
	}

	return createFoodCardUseCaseDto.CreateFoodCardUseCaseResponse{
		Code:     coreUseCaseResponseCode.SUCCESS,
		FoodCard: foodCard,
	}
}
