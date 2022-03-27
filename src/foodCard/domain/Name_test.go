package foodCardDomain

import (
	"github.com/bxcodec/faker/v3"
	"testing"
)

func TestNameCreate(t *testing.T) {
	nameString := faker.Name()
	name, err := CreateName(nameString)

	if err != nil || name.Value() != nameString {
		t.Errorf("CreateName(%s) should return Name{%s}, but got error: %v", nameString, nameString, err)
	}
}

func TestNameCreateFailWithEmptyNameString(t *testing.T) {
	nameString := ""
	_, err := CreateName(nameString)

	if err == nil {
		t.Errorf("nameString should not be an empty string, but no error occured")
	}
}
