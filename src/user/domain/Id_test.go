package userDomain

import (
	"testing"

	"github.com/bxcodec/faker/v3"
)

func TestCreateIdCreate(t *testing.T) {
	idString := faker.Email()
	id, err := CreateId(idString)

	if err != nil || id.Value() != idString {
		t.Errorf("CreateId(%s) should return Id{%s}, but got error: %v", idString, idString, err)
	}
}

func TestIdCreateFailWithEmptyIdString(t *testing.T) {
	idString := ""
	_, err := CreateId(idString)

	if err == nil {
		t.Errorf("idString should not be an empty string, but no error occured")
	}
}

func TestIdCreateFailWithInvalidIdString(t *testing.T) {
	idString := "abc@@gmail"
	_, err := CreateId(idString)

	if err == nil {
		t.Errorf("idString should not be an invalid email address, but no error occured")
	}
}
