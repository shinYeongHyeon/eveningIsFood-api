package foodCardCommandController

import (
	"github.com/gofiber/fiber/v2"
	createFoodCardUseCase "github.com/shinYeongHyeon/eveningIsFood-api/src/foodCard/application/CreateFoodCardUseCase"
	createFoodCardUseCaseDto "github.com/shinYeongHyeon/eveningIsFood-api/src/foodCard/application/CreateFoodCardUseCase/dto"
	foodCardControllerCommandDto "github.com/shinYeongHyeon/eveningIsFood-api/src/foodCard/controller/command/dto"
	foodCardDomain "github.com/shinYeongHyeon/eveningIsFood-api/src/foodCard/domain"
	postgresFoodCardRepository "github.com/shinYeongHyeon/eveningIsFood-api/src/foodCard/repository/postgres"
	coreUseCaseResponseCode "github.com/shinYeongHyeon/eveningIsFood-api/src/shared/core/useCase"
)

func Create(context *fiber.Ctx) error {
	request := new(foodCardControllerCommandDto.CreateRequest)

	if err := context.BodyParser(request); err != nil {
		return err
	}

	name, nameError := foodCardDomain.CreateName(request.Name)
	address, addressError := foodCardDomain.CreateAddress(request.Address)

	if nameError != nil || addressError != nil {
		return context.Status(fiber.StatusBadRequest).JSON(foodCardControllerCommandDto.CreateResponseErrorForBadRequest)
	}

	createFoodCardUseCaseResponse := createFoodCardUseCase.Execute(
		createFoodCardUseCaseDto.CreateFoodCardUseCaseRequest{
			Name:    name,
			Address: address,
		},
		postgresFoodCardRepository.GetRepository(),
	)

	if createFoodCardUseCaseResponse.Code != coreUseCaseResponseCode.SUCCESS {
		return context.Status(fiber.StatusInternalServerError).JSON(foodCardControllerCommandDto.CreateResponse{Code: "FAILED"})
	}

	foodCard := createFoodCardUseCaseResponse.FoodCard

	return context.Status(fiber.StatusCreated).JSON(foodCardControllerCommandDto.CreateResponse{
		Code: "SUCCESS",
		FoodCard: foodCardControllerCommandDto.FoodCard{
			Uuid:    foodCard.GetUuid(),
			Name:    foodCard.GetName().Value(),
			Address: foodCard.GetAddress().Value(),
		},
	})
}
