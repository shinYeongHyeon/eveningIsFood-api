package foodCardDomain

import coreUuid "github.com/shinYeongHyeon/eveningIsFoodApi/src/shared/core/uuid"

// FoodCard is domain type definition.
type FoodCard struct {
	Uuid    string
	Name    Name
	Address Address
}

// NewFoodCardProps is a struct for creating a new FoodCard domain.
type NewFoodCardProps struct {
	Name    Name
	Address Address
}

// CreateNewFoodCard creates a new FoodCard domain.
func CreateNewFoodCard(props NewFoodCardProps) (FoodCard, error) {
	return FoodCard{
		Uuid:    coreUuid.CreateUuid(),
		Name:    props.Name,
		Address: props.Address,
	}, nil
}

// GetUuid returns the Uuid.
func (foodCard FoodCard) GetUuid() string {
	return foodCard.Uuid
}

// GetName returns the Name domain.
func (foodCard FoodCard) GetName() Name {
	return foodCard.Name
}

// GetAddress returns the Address domain.
func (foodCard FoodCard) GetAddress() Address {
	return foodCard.Address
}
