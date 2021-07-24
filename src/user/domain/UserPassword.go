package user_domain

import (
"errors"
"unicode/utf8"
)

// UserPassword structure
type UserPassword struct {
	Value string
}

// UserPasswordCreate create UserPassword with userPasswordString
func UserPasswordCreate(userPasswordString string) (UserPassword, error) {
	if userPasswordString == "" {
		return UserPassword{}, errors.New("userPasswordString Can Not Be Empty")
	}

	if utf8.RuneCountInString(userPasswordString) < 6 {
		return UserPassword{}, errors.New("userPasswordString Length Can Not Be Less Than 6")
	}

	return UserPassword { userPasswordString }, nil
}