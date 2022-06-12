package foodShopCardDomain

import coreUuid "github.com/shinYeongHyeon/eveningIsFood-api/src/shared/core/uuid"

// FoodShopCard is domain type definition.
type FoodShopCard struct {
	uuid    string
	name    Name
	address Address
}

// NewFoodShopCardProps is a struct for creating a new FoodShopCard domain.
type NewFoodShopCardProps struct {
	Name    Name
	Address Address
}

// CreateNewFoodShopCard creates a new FoodShopCard domain.
func CreateNewFoodShopCard(props NewFoodShopCardProps) (FoodShopCard, error) {
	return FoodShopCard{
		uuid:    coreUuid.CreateUuid(),
		name:    props.Name,
		address: props.Address,
	}, nil
}

// GetUuid returns the Uuid.
func (foodShopCard FoodShopCard) GetUuid() string {
	return foodShopCard.uuid
}

// GetName returns the Name domain.
func (foodShopCard FoodShopCard) GetName() Name {
	return foodShopCard.name
}

// GetAddress returns the Address domain.
func (foodShopCard FoodShopCard) GetAddress() Address {
	return foodShopCard.address
}
