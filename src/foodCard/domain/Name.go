package foodCardDomain

import "errors"

// Name is value object type definition.
type Name struct {
	value string
}

// CreateName creates a Name domain.
func CreateName(nameString string) (Name, error) {
	if nameString == "" {
		return Name{}, errors.New("nameString is empty")
	}

	return Name{nameString}, nil
}

// Value returns the value of the Name domain.
func (name Name) Value() string {
	return name.value
}
