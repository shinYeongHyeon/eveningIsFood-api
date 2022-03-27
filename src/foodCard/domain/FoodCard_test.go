package foodCardDomain

import (
	"github.com/bxcodec/faker/v3"
	"testing"
)

func TestNewFoodCardCreate(t *testing.T) {
	nameString := faker.Name()
	addressString := faker.Sentence()

	foodCardProps := NewFoodCardProps{
		Name:    Name{nameString},
		Address: Address{addressString},
	}
	foodCard, err := CreateNewFoodCard(foodCardProps)

	if err != nil || foodCard.GetName().Value() != nameString || foodCard.GetAddress().Value() != addressString {
		t.Errorf("CreateNewFoodCard(%s, %s) should return FoodCard{%s, %s}, but got error: %v", nameString, addressString, nameString, addressString, err)
	}
}
