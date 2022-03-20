package userDomain

import (
	"testing"

	"github.com/bxcodec/faker/v3"
)

func TestNewUserCreate(t *testing.T) {
	nameString := faker.Name()
	idString := faker.Email()
	passwordString := "superStrongPassword"

	userProps := NewUserProps{
		name:     Name{nameString},
		id:       Id{idString},
		password: Password{passwordString},
	}

	user, err := CreateNewUser(userProps)

	if err != nil || user.GetName().Value() != nameString && user.GetId().Value() != idString || user.GetPassword().Value() != passwordString {
		t.Errorf("CreateNewUser(%v) should return User{%s, %s, %s}, but got error: %v", userProps, nameString, idString, passwordString, err)
	}
}
