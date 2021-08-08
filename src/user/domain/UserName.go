package user_domain

import (
	"errors"
	"unicode/utf8"
)

// UserName structure
type UserName struct {
	value string
}

// UserNameCreate create UserName with userNameString
func UserNameCreate(userNameString string) (UserName, error) {
	if userNameString == "" {
		return UserName{}, errors.New("userNameString Can Not Be Empty")
	}

	if utf8.RuneCountInString(userNameString) < 3 {
		return UserName{}, errors.New("userNameString Length Can Not Be Less Than 3")
	}

	return UserName{userNameString}, nil
}

func (userName *UserName) Value() string {
	return userName.value
}
