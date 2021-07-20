package user_domain

import "errors"

// UserName structure
type UserName struct {
	Value string
}

// UserNameCreate create UserName with userNameString
func UserNameCreate(userNameString string) (UserName, error) {
	if userNameString == "" {
		return UserName{}, errors.New("userNameString Can Not Be Empty")
	}

	return UserName { userNameString }, nil
}