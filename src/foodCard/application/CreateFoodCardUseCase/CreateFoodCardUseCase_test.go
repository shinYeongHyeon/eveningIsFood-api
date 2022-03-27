package createFoodCardUseCase

import (
	"github.com/bxcodec/faker/v3"
	createFoodCardUseCaseDto "github.com/shinYeongHyeon/eveningIsFood-api/src/foodCard/application/CreateFoodCardUseCase/dto"
	foodCardDomain "github.com/shinYeongHyeon/eveningIsFood-api/src/foodCard/domain"
	testFoodCardRepository "github.com/shinYeongHyeon/eveningIsFood-api/src/foodCard/repository/test"
	coreUseCaseResponseCode "github.com/shinYeongHyeon/eveningIsFood-api/src/shared/core/useCase"
	"testing"
)

func TestExecute(t *testing.T) {
	fakerName := faker.Name()
	fakerSentence := faker.Sentence()
	Name, _ := foodCardDomain.CreateName(fakerName)
	Address, _ := foodCardDomain.CreateAddress(fakerSentence)

	response := Execute(
		createFoodCardUseCaseDto.CreateFoodCardUseCaseRequest{
			Name:    Name,
			Address: Address,
		},
		testFoodCardRepository.GetRepository(false),
	)

	if response.Code != coreUseCaseResponseCode.SUCCESS {
		t.Errorf("Expected code : %s, but got code : %s", coreUseCaseResponseCode.SUCCESS, response.Code)
	}
}

func TestSaveFailedExecute(t *testing.T) {
	Name, _ := foodCardDomain.CreateName(faker.Name())
	Address, _ := foodCardDomain.CreateAddress(faker.Sentence())

	response := Execute(
		createFoodCardUseCaseDto.CreateFoodCardUseCaseRequest{
			Name:    Name,
			Address: Address,
		},
		testFoodCardRepository.GetRepository(true),
	)

	if response.Code != createFoodCardUseCaseDto.CodeFailedSaveFoodCardDomain {
		t.Errorf("Expected code : %s, but got code : %s", createFoodCardUseCaseDto.CodeFailedSaveFoodCardDomain, response.Code)
	}
}
