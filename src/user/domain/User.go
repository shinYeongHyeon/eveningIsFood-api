package user_domain

import (
	core_domain "github.com/shinYeongHyeon/eveningIsFood/eveningIsFoodApi/core/domain"
)

type User struct {
	Uuid string
	Email UserEmail
	Name UserName
	Password UserPassword
}

type NewUserProps struct {
	UserEmail UserEmail
	UserName UserName
	UserPassword UserPassword
}

func NewUserCreate(props NewUserProps) (User, error) {
	return User {
		Uuid: core_domain.CreateUuid(),
		Email: props.UserEmail,
		Name: props.UserName,
		Password: props.UserPassword,
	}, nil
}