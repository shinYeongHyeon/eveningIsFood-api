package foodShopCardCommandController

import (
	"github.com/gofiber/fiber/v2"
	createFoodShopCardUseCase "github.com/shinYeongHyeon/eveningIsFood-api/src/foodShopCard/application/CreateFoodShopCardUseCase"
	createFoodShopCardUseCaseDto "github.com/shinYeongHyeon/eveningIsFood-api/src/foodShopCard/application/CreateFoodShopCardUseCase/dto"
	foodShopCardControllerCommandDto "github.com/shinYeongHyeon/eveningIsFood-api/src/foodShopCard/controller/command/dto"
	foodShopCardDomain "github.com/shinYeongHyeon/eveningIsFood-api/src/foodShopCard/domain"
	postgresFoodShopCardRepository "github.com/shinYeongHyeon/eveningIsFood-api/src/foodShopCard/repository/postgres"
	coreUseCaseResponseCode "github.com/shinYeongHyeon/eveningIsFood-api/src/shared/core/useCase"
)

// Create : Create New foodShopCard
// @Summary      음식점카드 생성
// @Description  음식점카드를 생성하는 API
// @Tags         foodShopCard
// @Accept       json
// @Produce      json
// @Param        foodShopCard body     foodShopCardControllerCommandDto.CreateRequest       true "음식점카드 생성 Request Body"
// @Success      200      {object} foodShopCardControllerCommandDto.CreateResponse      "음식점카드 생성 Response"
// @Failure      400      {object} foodShopCardControllerCommandDto.CreateResponseError
// @Failure      500      {object} foodShopCardControllerCommandDto.CreateResponseError
// @Router       /foodShopCard [post]
func Create(context *fiber.Ctx) error {
	request := new(foodShopCardControllerCommandDto.CreateRequest)

	if err := context.BodyParser(request); err != nil {
		return err
	}

	name, nameError := foodShopCardDomain.CreateName(request.Name)
	address, addressError := foodShopCardDomain.CreateAddress(request.Address)

	if nameError != nil || addressError != nil {
		return context.Status(fiber.StatusBadRequest).JSON(foodShopCardControllerCommandDto.CreateResponseErrorForBadRequest)
	}

	createFoodShopCardUseCaseResponse := createFoodShopCardUseCase.Execute(
		createFoodShopCardUseCaseDto.CreateFoodShopCardUseCaseRequest{
			Name:    name,
			Address: address,
		},
		postgresFoodShopCardRepository.GetRepository(),
	)

	if createFoodShopCardUseCaseResponse.Code != coreUseCaseResponseCode.SUCCESS {
		return context.Status(fiber.StatusInternalServerError).JSON(foodShopCardControllerCommandDto.CreateResponseError{
			Code:         "FAILED",
			ErrorMessage: "Failed to create foodShopCard.",
		})
	}

	foodShopCard := createFoodShopCardUseCaseResponse.FoodShopCard

	return context.Status(fiber.StatusCreated).JSON(foodShopCardControllerCommandDto.CreateResponse{
		Code: "SUCCESS",
		FoodShopCard: foodShopCardControllerCommandDto.FoodShopCard{
			Uuid:    foodShopCard.GetUuid(),
			Name:    foodShopCard.GetName().Value(),
			Address: foodShopCard.GetAddress().Value(),
		},
	})
}
