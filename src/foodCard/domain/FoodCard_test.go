package foodCardDomain

import "testing"

func TestNewFoodCardCreate(t *testing.T) {
	nameString := "초원"
	addressString := "서울특별시 강남구 테헤란로2길 21"

	foodCardProps := NewFoodCardProps{
		Name:    Name{nameString},
		Address: Address{addressString},
	}
	foodCard, err := CreateNewFoodCard(foodCardProps)

	if err != nil || foodCard.GetName().Value() != nameString || foodCard.GetAddress().Value() != addressString {
		t.Errorf("CreateNewFoodCard(%s, %s) should return FoodCard{%s, %s}, but got error: %v", nameString, addressString, nameString, addressString, err)
	}
}
