package foodCardDomain

import "testing"

func TestAddressCreate(t *testing.T) {
	addressString := "서울특별시 강남구 테헤란로2길 21"
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
