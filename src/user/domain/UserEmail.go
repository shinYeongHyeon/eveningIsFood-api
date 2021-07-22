package user_domain

import (
	"errors"
	"regexp"
)

// UserEmail structure
type UserEmail struct {
	Value string
}

// UserEmailCreate create UserEmail with userEmailString
func UserEmailCreate(userEmailString string) (UserEmail, error) {
	if !regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`).MatchString(userEmailString) {
		return UserEmail{}, errors.New("userEmailString Must be email form")
	}

	return UserEmail { userEmailString }, nil
}