package userdomain

import (
	core_domain "github.com/shinYeongHyeon/eveningIsFood/eveningIsFoodApi/core/domain"
)

// User domain type
type User struct {
	Uuid     string
	Email    UserEmail
	Name     UserName
	Password UserPassword
}

// NewUserProps domain type for new one
type NewUserProps struct {
	UserEmail    UserEmail
	UserName     UserName
	UserPassword UserPassword
}

// NewUserCreate create new user
func NewUserCreate(props NewUserProps) (User, error) {
	return User{
		Uuid:     core_domain.CreateUuid(),
		Email:    props.UserEmail,
		Name:     props.UserName,
		Password: props.UserPassword,
	}, nil
}
