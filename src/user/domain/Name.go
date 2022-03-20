package userDomain

import "errors"

// Name is value object type definition.
type Name struct {
	value string
}

// CreateName creates a new name
func CreateName(nameString string) (Name, error) {
	if nameString == "" {
		return Name{}, errors.New("nameString is empty")
	}

	return Name{nameString}, nil
}

// Value returns the name
func (name Name) Value() string {
	return name.value
}
