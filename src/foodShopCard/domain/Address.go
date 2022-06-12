package foodShopCardDomain

import "errors"

// Address is value object type definition.
type Address struct {
	value string
}

// CreateAddress creates an Address domain.
func CreateAddress(addressString string) (Address, error) {
	if addressString == "" {
		return Address{}, errors.New("addressString is empty")
	}

	return Address{addressString}, nil
}

// Value returns the value of the Address domain.
func (address Address) Value() string {
	return address.value
}
