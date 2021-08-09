package createuserusecase

import (
	"fmt"
	createuserusecasedto "github.com/shinYeongHyeon/eveningIsFood/eveningIsFoodApi/src/user/application/CreateUserUseCase/dto"
)

// Exec Execute create user
func Exec(request createuserusecasedto.Request) {
	fmt.Println(request)
}
