package createuserusecase

import (
	createuserusecasedto "github.com/shinYeongHyeon/eveningIsFood/eveningIsFoodApi/src/user/application/CreateUserUseCase/dto"
	"testing"
)

// TestFailToCreateUserUseCase Testing for CreateUserUseCase Fail
func TestFailToCreateUserUseCase(t *testing.T) {
	isSuccess, err := Exec(createuserusecasedto.Request{
		Email:    "",
		Name:     "",
		Password: "",
	})

	if isSuccess || err == nil {
		t.Fatal("Fail to TestFailToCreateUserUseCase")
	}
}

// TestCreateUserUseCase Testing for CreateUserUseCase Success
func TestCreateUserUseCase(t *testing.T) {
	isSuccess, err := Exec(createuserusecasedto.Request{
		Email:    "den.shin.dev@gmail.com",
		Name:     "신영현",
		Password: "비밀번호123",
	})

	if !isSuccess || err != nil {
		t.Fatal("Fail to TestCreateUserUseCase")
	}
}
