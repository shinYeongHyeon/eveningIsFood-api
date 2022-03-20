package userDomain

import "errors"

// Password is value object type definition.
type Password struct {
	value string
}

// CreatePassword creates a new password
func CreatePassword(passwordString string) (Password, error) {
	if passwordString == "" {
		return Password{}, errors.New("passwordString is empty")
	}

	if len(passwordString) < 8 {
		return Password{}, errors.New("passwordString is too short")
	}

	return Password{passwordString}, nil
}

// Value returns the password
func (password Password) Value() string {
	return password.value
}
