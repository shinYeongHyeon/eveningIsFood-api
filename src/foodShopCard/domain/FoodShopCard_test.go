package foodShopCardDomain

import (
	"github.com/bxcodec/faker/v3"
	"testing"
)

func TestNewFoodShopCardCreate(t *testing.T) {
	nameString := faker.Name()
	addressString := faker.Sentence()

	foodShopCardProps := NewFoodShopCardProps{
		Name:    Name{nameString},
		Address: Address{addressString},
	}
	foodShopCard, err := CreateNewFoodShopCard(foodShopCardProps)

	if err != nil || foodShopCard.GetName().Value() != nameString || foodShopCard.GetAddress().Value() != addressString {
		t.Errorf("CreateNewFoodShopCard(%s, %s) should return FoodShopCard{%s, %s}, but got error: %v", nameString, addressString, nameString, addressString, err)
	}
}
