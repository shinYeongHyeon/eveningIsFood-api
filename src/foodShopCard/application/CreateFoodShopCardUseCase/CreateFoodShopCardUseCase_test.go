package createFoodShopCardUseCase

import (
	"github.com/bxcodec/faker/v3"
	createFoodShopCardUseCaseDto "github.com/shinYeongHyeon/eveningIsFood-api/src/foodShopCard/application/CreateFoodShopCardUseCase/dto"
	foodShopCardDomain "github.com/shinYeongHyeon/eveningIsFood-api/src/foodShopCard/domain"
	testFoodShopCardRepository "github.com/shinYeongHyeon/eveningIsFood-api/src/foodShopCard/repository/test"
	coreUseCaseResponseCode "github.com/shinYeongHyeon/eveningIsFood-api/src/shared/core/useCase"
	"testing"
)

func TestExecute(t *testing.T) {
	fakerName := faker.Name()
	fakerSentence := faker.Sentence()
	Name, _ := foodShopCardDomain.CreateName(fakerName)
	Address, _ := foodShopCardDomain.CreateAddress(fakerSentence)

	response := Execute(
		createFoodShopCardUseCaseDto.CreateFoodShopCardUseCaseRequest{
			Name:    Name,
			Address: Address,
		},
		testFoodShopCardRepository.GetRepository(false),
	)

	if response.Code != coreUseCaseResponseCode.SUCCESS {
		t.Errorf("Expected code : %s, but got code : %s", coreUseCaseResponseCode.SUCCESS, response.Code)
	}
}

func TestSaveFailedExecute(t *testing.T) {
	Name, _ := foodShopCardDomain.CreateName(faker.Name())
	Address, _ := foodShopCardDomain.CreateAddress(faker.Sentence())

	response := Execute(
		createFoodShopCardUseCaseDto.CreateFoodShopCardUseCaseRequest{
			Name:    Name,
			Address: Address,
		},
		testFoodShopCardRepository.GetRepository(true),
	)

	if response.Code != createFoodShopCardUseCaseDto.CodeFailedSaveFoodShopCardDomain {
		t.Errorf("Expected code : %s, but got code : %s", createFoodShopCardUseCaseDto.CodeFailedSaveFoodShopCardDomain, response.Code)
	}
}
