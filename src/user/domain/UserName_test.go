package user_domain

import (
	"testing"
)

// TestCreateUserName Testing for UserName Create Successful
func TestCreateUserName(t *testing.T) {
	userNameString := "Shin"
	_, err := UserNameCreate(userNameString)

	if err != nil {
		t.Fatal("Fail to Create UserName")
	}
}

// TestFailToCreateUserNameWhenUserNameStringIsEmpty Testing for UserName Create Fail When userNameString was Empty.
func TestFailToCreateUserNameWhenUserNameStringIsEmpty(t *testing.T) {
	userNameString := ""
	_, err := UserNameCreate(userNameString)

	if err == nil {
		t.Fatal("Fail to Create UserName")
	}
}