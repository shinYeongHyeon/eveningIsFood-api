package foodShopCardDomain

import (
	"github.com/bxcodec/faker/v3"
	"testing"
)

func TestAddressCreate(t *testing.T) {
	addressString := faker.Sentence()
	address, err := CreateAddress(addressString)

	if err != nil || address.Value() != addressString {
		t.Errorf("CreateAddress(%s) should return Address{%s}, but got error: %v", addressString, addressString, err)
	}
}

func TestAddressCreateFailWithEmptyNameString(t *testing.T) {
	addressString := ""
	_, err := CreateAddress(addressString)

	if err == nil {
		t.Errorf("addressString should not be an empty string, but no error occured")
	}
}
