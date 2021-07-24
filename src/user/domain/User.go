package user_domain

import "time"

type NewUser struct {
	Uuid string
	Email UserEmail
	Name UserName
	Password UserPassword
	CreatedAt time.Time
}

type NewUserProps struct {
	UserEmail UserEmail
	UserName UserName
	UserPassword UserPassword
}

func NewUserCreate(props NewUserProps) (NewUser, error) {
	return NewUser{}, nil
}