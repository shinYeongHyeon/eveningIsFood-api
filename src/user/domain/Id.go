package userDomain

import (
	"errors"
	"net/mail"
)

// Id is value object type definition.
type Id struct {
	value string
}

// CreateId creates a new email Id
func CreateId(idString string) (Id, error) {
	if idString == "" {
		return Id{}, errors.New("idString is empty")
	}

	if !validEmailForm(idString) {
		return Id{}, errors.New("idString is invalid email address")
	}

	return Id{idString}, nil
}

// Value returns the email address
func (id Id) Value() string {
	return id.value
}

func validEmailForm(emailAddress string) bool {
	_, err := mail.ParseAddress(emailAddress)
	return err == nil
}
