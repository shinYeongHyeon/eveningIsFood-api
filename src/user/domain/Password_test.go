package userDomain

import (
	"testing"
)

func TestPasswordCreate(t *testing.T) {
	passwordString := "superStrongPassowrd"
	password, err := CreatePassword(passwordString)

	if err != nil || password.Value() != passwordString {
		t.Errorf("CreatePassword(%s) should return Password{%s}, but got error: %v", passwordString, passwordString, err)
	}
}

func TestPasswordCreateWithEmptyPasswordString(t *testing.T) {
	passwordString := ""
	_, err := CreatePassword(passwordString)

	if err == nil {
		t.Errorf("passwordString should not be an empty string, but no error occured")
	}
}

func TestPasswordCreateWithTooShortPasswordString(t *testing.T) {
	passwordString := "easy"
	_, err := CreatePassword(passwordString)

	if err == nil {
		t.Errorf("passwordString should not be too short, but no error occured")
	}
}
