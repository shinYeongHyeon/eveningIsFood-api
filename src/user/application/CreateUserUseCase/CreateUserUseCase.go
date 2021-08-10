package createuserusecase

import (
	"errors"
	createuserusecasedto "github.com/shinYeongHyeon/eveningIsFood/eveningIsFoodApi/src/user/application/CreateUserUseCase/dto"
	userdomain "github.com/shinYeongHyeon/eveningIsFood/eveningIsFoodApi/src/user/domain"
)

// Exec Execute create user
func Exec(request createuserusecasedto.Request) (bool, error) {
	_, userNameErr := userdomain.UserNameCreate(request.Name)
	_, userEmailErr := userdomain.UserEmailCreate(request.Email)
	_, userPasswordErr := userdomain.UserPasswordCreate(request.Password)

	if userNameErr != nil || userEmailErr != nil || userPasswordErr != nil {
		return false, errors.New("err")
	}

	// TODO: Save And E2E Test

	return true, nil
}
