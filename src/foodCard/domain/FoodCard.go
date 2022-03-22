package foodCardDomain

import coreUuid "github.com/shinYeongHyeon/eveningIsFood-api/src/shared/core/uuid"

// FoodCard is domain type definition.
type FoodCard struct {
	uuid    string
	name    Name
	address Address
}

// NewFoodCardProps is a struct for creating a new FoodCard domain.
type NewFoodCardProps struct {
	Name    Name
	Address Address
}

// CreateNewFoodCard creates a new FoodCard domain.
func CreateNewFoodCard(props NewFoodCardProps) (FoodCard, error) {
	return FoodCard{
		uuid:    coreUuid.CreateUuid(),
		name:    props.Name,
		address: props.Address,
	}, nil
}

// GetUuid returns the Uuid.
func (foodCard FoodCard) GetUuid() string {
	return foodCard.uuid
}

// GetName returns the Name domain.
func (foodCard FoodCard) GetName() Name {
	return foodCard.name
}

// GetAddress returns the Address domain.
func (foodCard FoodCard) GetAddress() Address {
	return foodCard.address
}
